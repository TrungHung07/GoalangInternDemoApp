package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	pb "DemoApp/api/helloworld/v1"
	ent "DemoApp/ent"
	class "DemoApp/ent/class"
	"DemoApp/ent/teacher"
	"DemoApp/internal/data"
	"DemoApp/internal/utils"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	// "DemoApp/internal/biz"
)

// ClassServiceService handles the class-related business logic at the service layer
type ClassServiceService struct {
	pb.UnimplementedClassServiceServer
	s             *data.Data
	log           *log.Helper
	historyHelper *data.HistoryHelper
}

// NewClassServiceService creates a new instance of ClassServiceService with the given data layer, logger, and history helper.
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

func (s *ClassServiceService) saveClassToCache(ctx context.Context, key string, reply *pb.GetClassReply, id int64) {
	if data, err := json.Marshal(reply); err == nil {
		if cacheErr := s.setCache(ctx, key, data); cacheErr != nil {
			s.log.Warnf("Failed to set cache for class ID %d: %v", id, cacheErr)
		}
	}
}

func (s *ClassServiceService) saveClassListToCache(ctx context.Context, key string, reply *pb.ListClassReply) {
	if data, err := json.Marshal(reply); err == nil {
		_ = s.setCache(ctx, key, data)
	}
}

func (s *ClassServiceService) buildCacheKey(id int64) string {
	return fmt.Sprintf("class:%d", id)
}

func (s *ClassServiceService) getClassListFromCache(ctx context.Context, key string) *pb.ListClassReply {
	val, err := s.getCache(ctx, key)
	if err == nil && len(val) > 0 {
		log.Infof("Cache hit")
		reply := &pb.ListClassReply{}
		if e := json.Unmarshal(val, reply); e == nil {
			return reply
		}
	}
	return nil
}

func (s *ClassServiceService) getClassFromCache(ctx context.Context, key string) (*pb.GetClassReply, bool) {
	val, err := s.getCache(ctx, key)
	if err != nil || len(val) == 0 {
		return nil, false
	}

	var reply pb.GetClassReply
	if err := json.Unmarshal(val, &reply); err != nil {
		s.log.Warnf("Failed to unmarshal cached class: %v", err)
		return nil, false
	}
	return &reply, true
}

func (s *ClassServiceService) invalidateClassListCache(ctx context.Context) {
	iter := s.s.Redis.Scan(ctx, 0, "class:list:*", 0).Iterator()
	for iter.Next(ctx) {
		_ = s.s.Redis.Del(ctx, iter.Val()).Err()
	}
}

// CreateClass adds a new class using the business layer.
// It returns a CreateClassReply containing the result for the client.
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

// UpdateClass update a existed class using the business layer
// It returns a UpdateClassReply containing the result for the client
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

// DeleteClass removes a class from the database using the provided class ID.
// It returns a DeleteClassReply indicating the result of the operation.
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

func applyClassFilters(ctx context.Context, query *ent.ClassQuery, req *pb.ListClassRequest) *ent.ClassQuery {
	filter := req.Filter
	if filter == nil {
		return query
	}
	if filter.Name != nil {
		query = query.Where(class.NameContains(*filter.Name))
	}
	if filter.IsDeleted != nil {
		query = query.Where(class.IsDeletedEQ(*filter.IsDeleted))
	}
	if filter.Keyword != nil {
		query = query.Where(
			class.NameContains(*filter.Keyword),
		)
	}
	if filter.MinClassTeacher != nil {
		query = query.Where(class.HasTeachersWith(
			teacher.AgeGT(int(*filter.MinClassTeacher)),
		))
	}
	// if filter.MaxClassStudentQuantity != nil {
	// 	query = query.
	// }
	if filter.MaxClassStudentQuantity != nil {
		n := int(*filter.MaxClassStudentQuantity)

		selectQuery := query.Clone().Select(class.FieldID).Modify(

			func(s *sql.Selector) {
				studentTable := sql.Table("students")
				s.Join(studentTable).On(
					s.C("id"), studentTable.C("class_id"),
				)
				s.GroupBy(s.C("id"))
				s.Having(sql.LTE(sql.Count(studentTable.C("id")), n))
			},
		)
		var ids []int
		// selectQuery = selectQuery.Debug()
		if e := selectQuery.Scan(ctx, &ids); e != nil {
			return query
		}

		if len(ids) > 0 {
			query = query.Where(class.IDIn(ids...))
		} else {
			query = query.Where(class.IDEQ(-1)) // để trả về rỗng
		}
	}

	return query
}

// ListClass get all classes from the database which match filters ,paginations
// It returns a ListClassReply containing the result for the client
func (s *ClassServiceService) ListClass(ctx context.Context, req *pb.ListClassRequest) (*pb.ListClassReply, error) {
	// max:=int32(0); minT:=int32(0) ;
	// if req.Filter.MaxClassStudentQuantity == nil {maxQ =0}
	// if req.Filter.MinClassTeacher != nil {minT = req.Filter.MinClassTeacher }
	cacheKey := fmt.Sprintf("class:list:%v:%v:%v ", req.Filter, req.Page, req.PageSize)
	if cached := s.getClassListFromCache(ctx, cacheKey); cached != nil {
		return cached, nil // ← return luôn nếu cache hit
	}

	query := s.s.DB.Class.Query()
	query = applyClassFilters(ctx, query, req)
	total, e := query.Clone().Count(ctx)
	if e != nil {
		return nil, e
	}

	query = applyPagination(query, int(req.Page), int(req.PageSize))
	classes, e := query.All(ctx)
	if e != nil {
		return nil, e
	}
	dtoList := s.toDTOList(classes)
	reply := &pb.ListClassReply{
		Items: dtoList,
		Total: int64(total), // hoặc query tổng nếu dùng pagination
	}
	s.saveClassListToCache(ctx, cacheKey, reply)
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

// GetClass geta class from the database with Id
// It returns a GetClassReply containing the result for the client
func (s *ClassServiceService) GetClass(ctx context.Context, req *pb.GetClassRequest) (*pb.GetClassReply, error) {
	cacheKey := s.buildCacheKey(req.Id) // lấy key
	if reply, ok := s.getClassFromCache(ctx, cacheKey); ok {
		s.log.Infof("Cache hit for class ID %d", req.Id)
		return reply, nil
	}
	query := s.s.DB.Class.Query().Where(class.IDEQ(int(req.Id)))
	class, err := query.WithStudents().WithTeachers().Only(ctx)
	if err != nil {
		s.log.Infof("Failed to get class:", err)
		return nil, err
	}
	reply := s.toDTO(class)
	s.saveClassToCache(ctx, cacheKey, reply, req.Id)
	return reply, nil
}

func (s *ClassServiceService) toDTO(entity *ent.Class) *pb.GetClassReply {
	var classDTO pb.ClassData
	_ = copier.Copy(&classDTO, entity)
	studentsDTOList := s.mapStudents(entity.Edges.Students)
	teachersDTOList := s.mapTeachers(entity.Edges.Teachers)
	reply := &pb.GetClassReply{
		Class:            &classDTO,
		Students:         studentsDTOList,
		Teachers:         teachersDTOList,
		StudentsQuantity: int32(len(studentsDTOList)),
		TeachersQuantity: int32(len(teachersDTOList)),
	}
	return reply
}

func (s *ClassServiceService) toDTOList(classes []*ent.Class) []*pb.ClassData {
	var items []*pb.ClassData
	for _, class := range classes {
		items = append(items, &pb.ClassData{
			Id:        int64(class.ID),
			Name:      class.Name,
			Grade:     class.Grade,
			IsDeleted: &class.IsDeleted,
		})
	}
	return items
}

func (s *ClassServiceService) mapTeachers(teachers []*ent.Teacher) []*pb.TeacherDataForClass {
	var result []*pb.TeacherDataForClass
	for _, t := range teachers {
		result = append(result, &pb.TeacherDataForClass{
			Name: t.Name,
		})
	}
	return result
}

func (s *ClassServiceService) mapStudents(students []*ent.Student) []*pb.StudentDataForClass {
	var result []*pb.StudentDataForClass
	for _, t := range students {
		result = append(result, &pb.StudentDataForClass{
			Name: t.Name,
		})
	}
	return result
}

// ExportClassExcel geta class from the database with Id
// It returns a ExportClassExcelReply containing the result for the client
func (s *ClassServiceService) ExportClassExcel(ctx context.Context, req *pb.ExportClassExcelRequest) (*pb.ExportClassExcelReply, error) {
	// Fetch class data
	classResp, err := s.GetClass(ctx, &pb.GetClassRequest{Id: req.Id})
	if err != nil {
		return nil, fmt.Errorf("failed to get class: %w", err)
	}

	class := classResp.Class
	students := classResp.Students
	teachers := classResp.Teachers

	// Prepare students and teachers name lists
	var studentNames []string
	for _, s := range students {
		studentNames = append(studentNames, s.Name)
	}

	var teacherNames []string
	for _, t := range teachers {
		teacherNames = append(teacherNames, t.Name)
	}

	// Build report entry
	reportEntry := utils.ExcelClassReport{
		STT:             1,
		Class:           class.Name,
		Students:        studentNames,
		Teachers:        teacherNames,
		StudentQuantity: int(classResp.StudentsQuantity),
		TeacherQuantity: int(classResp.TeachersQuantity),
	}

	// Create Excel file
	helper := utils.NewExcelHelper()
	fileBytes, err := helper.ExportClassTableReport("BÁO CÁO ABC", []utils.ExcelClassReport{reportEntry})
	if err != nil {
		return nil, fmt.Errorf("failed to export excel: %w", err)
	}

	// Return as reply
	return &pb.ExportClassExcelReply{
		File: fileBytes,
	}, nil
}

func (s *ClassServiceService) ListExportClassExcelData(ctx context.Context, req *pb.ListClassExcelReportRequest) (*pb.ListClassExcelReportDataReply, error) {
	key := "list_report_data"
	var reply *pb.ListClassExcelReportDataReply
	val, err := s.getCache(ctx, key)
	if err == nil && len(val) > 0 {
		log.Infof("Cache hit")
		result := &pb.ListClassExcelReportDataReply{}
		if e := json.Unmarshal(val, reply); e == nil {
			reply = result
			return reply, nil
		}
	}
	query := s.s.DB.Class.Query().WithStudents().WithTeachers()
	entities, e := query.All(ctx)
	if e != nil {
		return nil, e
	}
	dataList := s.toReportDTOList(entities)
	reply = &pb.ListClassExcelReportDataReply{
		Data: dataList,
	}
	go func() {
		if data, err := json.Marshal(reply); err == nil {
			_ = s.setCache(ctx, key, data)
		}
	}()
	return reply, nil
}

func (s *ClassServiceService) toReportDTOList(classes []*ent.Class) []*pb.ReportExcelClassData {
	var result []*pb.ReportExcelClassData
	for _, c := range classes {
		result = append(result, s.toReportDTO(c))
	}
	return result
}

func (s *ClassServiceService) toReportDTO(class *ent.Class) *pb.ReportExcelClassData {
	students := s.mapStudentsForExport(class.Edges.Students)
	teachers := s.mapTeachersForExport(class.Edges.Teachers)

	return &pb.ReportExcelClassData{
		Class: &pb.ClassData{
			Id:    int64(class.ID),
			Name:  class.Name,
			Grade: class.Grade,
		},
		Students:         students,
		Teachers:         teachers,
		StudentsQuantity: int32(len(students)),
		TeachersQuantity: int32(len(teachers)),
	}
}

func (s *ClassServiceService) mapStudentsForExport(entities []*ent.Student) []*pb.ExportStudentDataForClass {
	var students []*pb.ExportStudentDataForClass
	for _, stu := range entities {
		students = append(students, &pb.ExportStudentDataForClass{
			Name: stu.Name,
		})
	}
	return students
}

func (s *ClassServiceService) mapTeachersForExport(entities []*ent.Teacher) []*pb.ExportTeacherDataForClass {
	var teachers []*pb.ExportTeacherDataForClass
	for _, t := range entities {
		teachers = append(teachers, &pb.ExportTeacherDataForClass{
			Name:  t.Name,
			Email: t.Email,
			Age:   int32(t.Age),
		})
	}
	return teachers
}

func (s *ClassServiceService) ExportListClassExcel(ctx context.Context, req *pb.ExportListClassExcelRequest) (*pb.ExportListClassExcelReply, error) {
	Data, e := s.ListExportClassExcelData(ctx, &pb.ListClassExcelReportRequest{})
	log.Info(Data)
	if e != nil {
		log.Errorf("ListExportClassExcelData failed: %v", e)
		return nil, e
	}
	var excelData []utils.ExcelClassData

	for _, c := range Data.Data {
		var students []utils.StudentInfo
		for _, stu := range c.Students {
			students = append(students, utils.StudentInfo{Name: stu.Name})
		}

		var teachers []utils.TeacherInfo
		for _, t := range c.Teachers {
			teachers = append(teachers, utils.TeacherInfo{
				Name:  t.Name,
				Email: t.Email,
				Age:   int(t.Age),
			})
		}

		excelData = append(excelData, utils.ExcelClassData{
			Class: utils.ClassInfo{
				ID:    int(c.Class.Id),
				Name:  c.Class.Name,
				Grade: int(c.Class.Grade),
			},
			Students:         students,
			Teachers:         teachers,
			StudentsQuantity: int(c.StudentsQuantity),
			TeachersQuantity: int(c.TeachersQuantity),
		})
	}

	helper := utils.NewExcelListHelper()
	log.Infof("Exporting %d classes to Excel", len(excelData))

	fileBytes, err := helper.ExportClassListExcel(excelData)
	if err != nil {
		return nil, err
	}

	return &pb.ExportListClassExcelReply{
		File: fileBytes,
	}, nil
}
