package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	// "fmt"

	pb "DemoApp/api/helloworld/v1"
	"DemoApp/internal/biz"
	"DemoApp/internal/common"
	"DemoApp/internal/common/filter"
	"DemoApp/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

//TeacherServiceService handles the teacher-related business logic at the service layer
type TeacherServiceService struct {
	pb.UnimplementedTeacherServiceServer
	data           *data.Data
	log            *log.Helper
	teacherUsecase *biz.TeacherUseCase
}

// NewTeacherServiceService creates a new instance of TeacherServiceService with the given data layer, use case logic, and logger.
func NewTeacherServiceService(data *data.Data, usecase *biz.TeacherUseCase, logger log.Logger) *TeacherServiceService {
	return &TeacherServiceService{
		data:           data,
		log:            log.NewHelper(logger),
		teacherUsecase: usecase,
	}
}


// GetTeacher retrieves a teacher by ID and returns the corresponding TeacherReply.
func (s *TeacherServiceService) GetTeacher(ctx context.Context, req *pb.GetTeacherRequest) (*pb.TeacherReply, error) {

	return &pb.TeacherReply{}, nil
}

// CreateTeacher adds a new teacher to the system based on the provided request data.
// It returns a CreateTeacherReply containing the newly created teacher's information.
func (s *TeacherServiceService) CreateTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*pb.CreateTeacherReply, error) {
	// teacher, err := s.data.DB.Teacher.Create().
	// 	SetName(req.Name).
	// 	SetEmail(req.Email).
	// 	SetClassID(int(req.ClassId)).
	// 	Save(ctx)
	// if err != nil {
	// 	s.log.Error("Failed to create teacher:", err)
	// 	return nil, err
	// }
	// s.log.Infof("Created teacher: %v", teacher)
	return &pb.CreateTeacherReply{Message: "Thêm thành công !"}, nil
}

// UpdateTeacher modifies the information of an existing teacher based on the request.
// It returns an UpdateTeacherReply with the updated data.
func (s *TeacherServiceService) UpdateTeacher(ctx context.Context, req *pb.UpdateTeacherRequest) (*pb.UpdateTeacherReply, error) {

	// teacher, err := s.data.DB.Teacher.Get(ctx, int(req.Id))
	// s.log.Infof("teacher and error :" + err.Error())
	// if err != nil {
	// 	s.log.Error("\nFailed to find teacher ( if is correct !):\n", err)
	// 	return &pb.UpdateTeacherReply{Message: "Không tìm thấy id!"}, nil
	// }
	// update := teacher.Update()
	// if req.Name != "" {
	// 	update.SetName(req.Name)
	// }
	// if req.Email != "" {
	// 	update.SetEmail(req.Email)
	// }
	// teacherUpdateed, err := update.Save(ctx)

	// if err != nil {
	// 	return &pb.UpdateTeacherReply{Message: "Cập nhật thất bại!"}, err
	// }
	// // teacher, err := s.data.DB.Teacher.UpdateOneID(int(req.Id)).
	// 	SetName(req.Name).
	// 	SetEmail(req.Email).
	// 	SetGrade(int(req.Grade)).
	// 	Save(ctx)
	// s.log.Infof("Updated teacher: %v", teacherUpdateed)
	return &pb.UpdateTeacherReply{Message: "Cập nhật thành công rồi đó !"}, nil
}

// DeleteTeacher removes a teacher from the system using the provided ID.
// It returns a DeleteTeacherReply indicating the result of the deletion.
func (s *TeacherServiceService) DeleteTeacher(ctx context.Context, req *pb.DeleteTeacherRequest) (*pb.DeleteTeacherReply, error) {

	// err := s.data.DB.Teacher.DeleteOneID(int(req.Id)).Exec(ctx)
	// if err != nil {
	// 	s.log.Error("Failed to delete teacher:", err)
	// 	return nil, err
	// }
	// s.log.Infof("Deleted teacher with ID: %d", req.Id)

	return &pb.DeleteTeacherReply{Message: "Xoá thành công !"}, nil
}

// ListTeachers returns a paginated list of teachers based on the request filters.
// It returns a ListTeachersReply containing the result set.
func (s *TeacherServiceService) ListTeachers(ctx context.Context, req *pb.ListTeachersRequest) (*pb.ListTeachersReply, error) {
	key := fmt.Sprintf("list teachers : %v %v %v", req.Page, req.PageSize, req.Fitler)
	val, err := s.data.RedisCache.Get(ctx, key)
	if err == nil && len(val) > 0 {
		s.log.Info("Cache hit for list teacher: ", key)
		reply := &pb.ListTeachersReply{}
		if e := json.Unmarshal(val, reply); e == nil {
			return reply, nil
		}
	}

	filter, pagination := convertListTeachersRequest(req)
	entites, err := s.teacherUsecase.FindAll(ctx, filter, pagination)
	if err != nil {
		s.log.Error("Failed to list teachers:", err)
		return nil, err
	}

	var teacherReplies []*pb.TeacherReply
	for _, entity := range entites {
		teacherReplies = append(teacherReplies, entity.To())
	}

	reply := &pb.ListTeachersReply{Teachers: teacherReplies}
	if data, e := json.Marshal(reply); e == nil {
		_ = s.data.RedisCache.Set(ctx, key, data, 10*time.Minute)
	}
	return reply, nil
}

func convertListTeachersRequest(req *pb.ListTeachersRequest) (filter.TeacherFilter, common.Pagination) {
	var f filter.TeacherFilter
	if req.Fitler != nil {
		f = filter.TeacherFilter{
			Email:  req.Fitler.Email,
			MinAge: req.Fitler.MinAge,
			MaxAge: req.Fitler.MaxAge,
		}
	}

	p := common.Pagination{
		Page:     int(req.GetPage()),
		PageSize: int(req.GetPageSize()),
	}

	return f, p
}
