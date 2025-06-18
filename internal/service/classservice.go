package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	pb "DemoApp/api/helloworld/v1"
	ent "DemoApp/ent"
	class "DemoApp/ent/class"
	"DemoApp/internal/data"

	"github.com/go-kratos/kratos/v2/log"
	// "DemoApp/internal/biz"
)

type ClassServiceService struct {
	pb.UnimplementedClassServiceServer
	s             *data.Data
	log           *log.Helper
	historyHelper *data.HistoryHelper
}

func NewClassServiceService(data *data.Data, logger log.Logger, helper *data.HistoryHelper) *ClassServiceService {
	return &ClassServiceService{
		s:             data,
		log:           log.NewHelper(logger),
		historyHelper: helper,
	}
}

// Lưu giá trị vào redis
// Nếu là struct, map, slice: Thường sẽ marshal thành JSON (json.Marshal) để ra []byte rồi lưu.
func (s *ClassServiceService) setCache(ctx context.Context, key string, value []byte) error {
	return s.s.Redis.Set(ctx, key, value, 10*time.Minute).Err()
}

func (s *ClassServiceService) getCache(ctx context.Context, key string) ([]byte, error) {
	return s.s.Redis.Get(ctx, key).Bytes()
}

func (s *ClassServiceService) invalidateClassListCache(ctx context.Context) {
	iter := s.s.Redis.Scan(ctx, 0, "class:list:*", 0).Iterator()
	for iter.Next(ctx) {
		_ = s.s.Redis.Del(ctx, iter.Val()).Err()
	}
}

func (s *ClassServiceService) CreateClass(ctx context.Context, req *pb.CreateClassRequest) (*pb.CreateClassReply, error) {
	//Tạo một entity
	classData := &ent.Class{
		Name:      req.Name,
		Grade:     req.Grade,
		IsDeleted: false,
	}
	//Lưu class vào database
	savedClass, err := s.s.DB.Class.Create().
		SetName(classData.Name).
		SetGrade(classData.Grade).
		SetIsDeleted(classData.IsDeleted).
		Save(ctx)
	if err != nil {
		s.log.Error("Failed to create class:", err)
		return nil, err
	}
	// Lưu lịch sử tạo lớp học
	_ = s.historyHelper.TrackInsert(ctx,
		"class",
		fmt.Sprint(savedClass.ID),
		savedClass,
		"system",
	)
	s.log.Info("Class created successfully:", classData.Name)
	s.invalidateClassListCache(ctx) // Xóa cache danh sách lớp học
	return &pb.CreateClassReply{Message: "Thêm lớp học thành công ! "}, nil
}

func (s *ClassServiceService) UpdateClass(ctx context.Context, req *pb.UpdateClassRequest) (*pb.UpdateClassReply, error) {
	classEntity, e := s.s.DB.Class.Get(ctx, int(req.Id))
	if e != nil {
		return nil, e
	}

	update := classEntity.Update()
	if req.Name != "" {
		update.SetName(req.Name)
	}

	if req.Grade != 0 {
		update.SetGrade(req.Grade)
	}

	_, e = update.Save(ctx)
	if e != nil {
		return nil, e
	}

	cacheKey := fmt.Sprintf("class:%d", req.Id)
	_ = s.s.Redis.Del(ctx, cacheKey).Err() // Xóa cache của lớp học đã cập nhật
	s.invalidateClassListCache(ctx)        // Xóa cache danh sách lớp học

	return &pb.UpdateClassReply{Message: "Cập nhật thành công !"}, nil
}

func (s *ClassServiceService) DeleteClass(ctx context.Context, req *pb.DeleteClassRequest) (*pb.DeleteClassReply, error) {
	classEntity, e := s.s.DB.Class.Get(ctx, int(req.Id))
	if e != nil {
		return nil, e
	}

	_, e = classEntity.Update().
		SetIsDeleted(true).
		Save(ctx)

	if e != nil {
		return nil, e
	}

	cacheKey := fmt.Sprintf("class:%d", req.Id)
	_ = s.s.Redis.Del(ctx, cacheKey).Err() // Xóa cache của lớp học đã cập nhật
	s.invalidateClassListCache(ctx)        // Xóa cache danh sách lớp học
	return &pb.DeleteClassReply{Message: "Xóa thành công lớp học này ! "}, nil
}

func applyPagination(query *ent.ClassQuery, page, pageSize int) *ent.ClassQuery {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return query.Offset(offset).Limit(pageSize)
}

func applyClassFilters(query *ent.ClassQuery, req *pb.ListClassRequest) *ent.ClassQuery {
	if req.Name != nil {
		query = query.Where(class.NameContains(*req.Name))
	}
	if req.Grade != nil {
		query = query.Where(class.GradeEQ(*req.Grade))
	}
	if req.IsDeleted != nil {
		query = query.Where(class.IsDeletedEQ(*req.IsDeleted))
	}
	if req.Keyword != nil {
		query = query.Where(
			class.NameContains(*req.Keyword),
		)
	}
	return query
}

func (s *ClassServiceService) ListClass(ctx context.Context, req *pb.ListClassRequest) (*pb.ListClassReply, error) {

	cacheKey := fmt.Sprintf("class:list:%v:%v:%v:%v:%v", req.Name, req.Grade, req.IsDeleted, req.Page, req.PageSize)
	val, err := s.getCache(ctx, cacheKey)
	if err == nil && len(val) > 0 {
		log.Infof("Cache hit")
		reply := &pb.ListClassReply{}
		if e := json.Unmarshal(val, reply); e == nil {
			return reply, nil
		}
	}

	query := s.s.DB.Class.Query()
	query = applyClassFilters(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		s.log.Infof("Failed to count classes:", err)
		return nil, err
	}
	query = applyPagination(query, int(req.Page), int(req.PageSize))
	classes, err := query.All(ctx)
	if err != nil {
		s.log.Infof("Failed to list classes:", err)
		return nil, err
	}
	var items []*pb.ClassData
	for _, class := range classes {
		item := &pb.ClassData{
			Id:        int64(class.ID),
			Name:      class.Name,
			Grade:     class.Grade,
			IsDeleted: &class.IsDeleted,
		}
		items = append(items, item)
	}
	reply := &pb.ListClassReply{
		Items: items,
		Total: int64(total), // hoặc query tổng nếu dùng pagination
	}

	if data, e := json.Marshal(reply); e == nil {
		_ = s.setCache(ctx, cacheKey, data)
	}

	return reply, nil
}

// redis lưu trữ key-value, key là class ID, value là dữ liệu class đã được serialize
// Nếu có cache, trả về dữ liệu từ cache
// Nếu không có cache, truy vấn từ database, sau đó lưu vào cache
// Nếu có lỗi trong quá trình lấy cache, tiếp tục truy vấn từ database
// Nếu có lỗi trong quá trình truy vấn database, trả về lỗi
// Nếu truy vấn thành công, lưu dữ liệu vào cache và trả về kết quả

// redis :
// key : string
// value : string,set,.... JSON
func (s *ClassServiceService) GetClass(ctx context.Context, req *pb.GetClassRequest) (*pb.GetClassReply, error) {
	cacheKey := fmt.Sprintf("class:%d", req.Id) // lấy key
	// Thử lấy cache từ Redis -> value
	val, err := s.getCache(ctx, cacheKey)

	if err == nil && len(val) > 0 {
		log.Infof("Cache hit for class ID: %d", req.Id)
		reply := &pb.GetClassReply{}
		//Unmarshal : chuyển đổi dữ liệu từ JSON về struct
		if unmarshalErr := json.Unmarshal(val, reply); unmarshalErr == nil {
			return reply, nil
		}
		// s.log.Warnf("Failed to unmarshal cached class: %v", unmarshalErr)
	}

	query := s.s.DB.Class.Query().Where(class.IDEQ(int(req.Id)))
	class, err := query.WithStudents().Only(ctx)
	if err != nil {
		s.log.Infof("Failed to get class:", err)
		return nil, err
	}
	var students []*pb.StudentDataForClass
	for _, student := range class.Edges.Students {
		students = append(students, &pb.StudentDataForClass{
			Id:   int64(student.ID),
			Name: student.Name,
		})
	}
	s.log.Infof("\nGet class with students: %v\n", students)
	// Convert class to pb.ClassData
	reply := &pb.GetClassReply{
		Class: &pb.ClassData{
			Id:        int64(class.ID),
			Name:      class.Name,
			Grade:     class.Grade,
			IsDeleted: &class.IsDeleted,
		},
		Students: students,
	}

	//Lưu vào cache
	//Marshal : chuyển đổi dữ liệu từ struct về JSON để lưu vào Redis cache : value
	if data, e := json.Marshal(reply); e == nil {
		if cacheErr := s.setCache(ctx, cacheKey, data); cacheErr != nil {
			s.log.Warnf("Failed to set cache for class ID %d: %v", req.Id, cacheErr)
		}
	}

	return reply, nil
}
