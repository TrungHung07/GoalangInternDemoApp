package service

import (
	pb "DemoApp/api/helloworld/v1"
	"DemoApp/ent"
	"DemoApp/ent/student"
	"DemoApp/internal/data"
	"encoding/json"
	"fmt"
	"time"

	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type StudentServiceService struct {
	pb.UnimplementedStudentServiceServer
	data          *data.Data
	log           *log.Helper
	hisotryHelper *data.HistoryHelper
}

func NewStudentServiceService(data *data.Data, logger log.Logger, helper *data.HistoryHelper) *StudentServiceService {
	return &StudentServiceService{
		data:          data,
		log:           log.NewHelper(logger),
		hisotryHelper: helper,
	}
}

func studentPagination(query *ent.StudentQuery, page, pageSize int) *ent.StudentQuery {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return query.Offset(offset).Limit(pageSize)
}

func (s *StudentServiceService) setCache(ctx context.Context, key string, value []byte) error {
	return s.data.Redis.Set(ctx, key, value, 5*time.Minute).Err()
}

func (s *StudentServiceService) getCache(ctx context.Context, key string) ([]byte, error) {
	return s.data.Redis.Get(ctx, key).Bytes()
}

func (s *StudentServiceService) invalidateStudentListCache(ctx context.Context) {
	iter := s.data.Redis.Scan(ctx, 0, "student:list:*", 0).Iterator()
	for iter.Next(ctx) {
		_ = s.data.Redis.Del(ctx, iter.Val()).Err()
	}
}

func (s *StudentServiceService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	studentData := &ent.Student{
		Name:      req.Name,
		ClassID:   int(req.ClassId),
		IsDeleted: false,
	}
	// Save the student to the database
	savedStudent, err := s.data.DB.Student.Create().
		SetName(studentData.Name).
		SetClassID(studentData.ClassID).
		SetIsDeleted(studentData.IsDeleted).
		Save(ctx)
	if err != nil {
		s.log.Error("Failed to create student:", err)
		return nil, err
	}
	_ = s.hisotryHelper.TrackInsert(ctx,
		"student",
		fmt.Sprint(savedStudent.ID),
		savedStudent,
		"system")
	s.log.Info("Student created successfully:", studentData.Name)
	s.invalidateStudentListCache(ctx) // Invalidate the student list cache
	return &pb.CreateStudentReply{Message: "Thêm thành công 1 học sinh mới "}, nil
}

func (s *StudentServiceService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {

	studentEntity, err := s.data.DB.Student.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	update := studentEntity.Update()
	if req.Name != "" {
		update.SetName(req.Name)
	}
	if req.ClassId != 0 {
		update.SetClassID(int(req.ClassId))
	}
	_, err = update.Save(ctx)

	if err != nil {
		return nil, err
	}
	cacheKey := fmt.Sprintf("student:%d", req.Id)
	_ = s.data.Redis.Del(ctx, cacheKey).Err() // Xóa cache của học sinh đã cập nhật
	s.invalidateStudentListCache(ctx)         // Xóa cache danh sách học sinh

	return &pb.UpdateStudentReply{Message: "Cập nhật thành công học sinh "}, nil
}

func (s *StudentServiceService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	err := s.data.DB.Student.UpdateOneID(int(req.Id)).SetIsDeleted(true).Exec(ctx)
	if err != nil {
		s.log.Error("Failed to delete student:", err)
		return nil, err
	}

	cacheKey := fmt.Sprintf("student:%d", req.Id)
	_ = s.data.Redis.Del(ctx, cacheKey).Err() // Xóa cache của học sinh đã xóa
	s.invalidateStudentListCache(ctx)         // Xóa cache danh sách học sinh
	return &pb.DeleteStudentReply{Message: "Xóa thành công học sinh này !"}, nil
}

func (s *StudentServiceService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {

	cacheKey := fmt.Sprintf("student:list:%d:%d", req.Page, req.PageSize)
	val, err := s.getCache(ctx, cacheKey)
	if err == nil && len(val) > 0 {
		s.log.Info("Cache hit for student list:", cacheKey)
		reply := &pb.ListStudentReply{}
		if e := json.Unmarshal(val, reply); e == nil {
			return reply, nil
		}
	}

	query := s.data.DB.Student.Query()
	// query = applyStudentFilters(query, req)
	query = studentPagination(query, int(req.Page), int(req.PageSize))
	total, err := query.Count(ctx)
	if err != nil {
		s.log.Error("Failed to count students:", err)
		return nil, err
	}

	students, err := query.WithClasses().All(ctx)
	if err != nil {
		s.log.Error("Failed to list students:", err)
		return nil, err
	}

	var items []*pb.StudentData

	for _, student := range students {
		className := ""
		if student.Edges.Classes != nil {
			className = student.Edges.Classes.Name
		}
		items = append(items, &pb.StudentData{
			Id:        int64(student.ID),
			Name:      student.Name,
			ClassId:   int64(student.ClassID),
			ClassName: className,
			IsDeleted: &student.IsDeleted,
		})
	}
	reply := &pb.ListStudentReply{
		Total: int64(total),
		Items: items,
	}

	if data, e := json.Marshal(reply); e == nil {
		_ = s.setCache(ctx, cacheKey, data)
	}

	return reply, nil
}

func (s *StudentServiceService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {

	cacheKey := fmt.Sprintf("student:%d", req.Id)
	val, err := s.getCache(ctx, cacheKey)
	if err == nil && len(val) > 0 {
		s.log.Info("Cache hit for student:", req.Id)
		reply := &pb.GetStudentReply{}
		if e := json.Unmarshal(val, reply); e == nil {
			return reply, nil
		}
	}

	query := s.data.DB.Student.Query().WithClasses().Where(student.IDEQ(int(req.Id)))
	student, err := query.Only(ctx)
	if err != nil {
		s.log.Error("Failed to get student:", err)
		return nil, err
	}

	var studentDTO *pb.StudentData
	className := ""
	if student.Edges.Classes != nil {
		className = student.Edges.Classes.Name
	}

	reply := &pb.GetStudentReply{
		Student: &pb.StudentData{
			Id:        int64(student.ID),
			Name:      student.Name,
			ClassId:   int64(student.ClassID),
			ClassName: className,
			IsDeleted: &student.IsDeleted,
		},
	}

	if data, e := json.Marshal(&pb.GetStudentReply{Student: studentDTO}); e == nil {
		if err := s.setCache(ctx, cacheKey, data); err != nil {
			s.log.Error("Failed to set cache for student:", err)
		}
	}

	return reply, nil
}
