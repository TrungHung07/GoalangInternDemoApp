package service

import (
	"context"
	"encoding/json"
	"errors"
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

// TeacherServiceService handles the teacher-related business logic at the service layer
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
	key := fmt.Sprintf("teacher: %v", req.Id)
	val, err := s.data.RedisCache.Get(ctx, key)
	if err == nil && len(val) > 0 {
		s.log.Info("Cache hit for teacher: ", key)
		reply := &pb.TeacherReply{}
		if e := json.Unmarshal(val, reply); e == nil {
			return reply, nil
		}
	}

	entity, e := s.teacherUsecase.FindByID(ctx, req.Id)
	if e != nil {
		return nil, e
	}

	reply := &pb.TeacherReply{
		Id:        entity.ID,
		Name:      entity.Name,
		Age:       entity.Age,
		Email:     entity.Email,
		ClassName: entity.ClassName,
	}

	go func() {
		if data, e := json.Marshal(reply); e == nil {
			_ = s.data.RedisCache.Set(ctx, key, data, 10*time.Minute)
		}
	}()

	return reply, nil
}

func (s *TeacherServiceService) invalidateTeacherListCache(ctx context.Context) {
	iter := s.data.Redis.Scan(ctx, 0, "list teachers : *", 0).Iterator()
	for iter.Next(ctx) {
		_ = s.data.Redis.Del(ctx, iter.Val()).Err()
	}
}

// CreateTeacher adds a new teacher to the system based on the provided request data.
// It returns a CreateTeacherReply containing the newly created teacher's information.
func (s *TeacherServiceService) CreateTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*pb.CreateTeacherReply, error) {
	teacher := &biz.Teacher{
		Name:    req.Name,
		Email:   req.Email,
		Age:     req.Age,
		ClassID: int64(req.ClassId),
	}
	_, e := s.teacherUsecase.Create(ctx, teacher)
	if e != nil {
		return nil, e
	}
	s.invalidateTeacherListCache(ctx)
	return &pb.CreateTeacherReply{Message: "Thêm thành công !"}, nil
}

// UpdateTeacher modifies the information of an existing teacher based on the request.
// It returns an UpdateTeacherReply with the updated data.
func (s *TeacherServiceService) UpdateTeacher(ctx context.Context, req *pb.UpdateTeacherRequest) (*pb.UpdateTeacherReply, error) {
	if req.Id == nil {
		return nil, errors.New("missing teacher ID for update")
	}

	teacher := &biz.UpdateTeacher{
		Name:    req.Name,
		Email:   req.Email,
		Age:     req.Age,
		ClassID: req.ClassId,
	}
	e := s.teacherUsecase.Update(ctx, teacher, *req.Id)
	if e != nil {
		return nil, e
	}
	return &pb.UpdateTeacherReply{Message: "Cập nhật thành công rồi đó !"}, nil
}

// DeleteTeacher removes a teacher from the system using the provided ID.
// It returns a DeleteTeacherReply indicating the result of the deletion.
func (s *TeacherServiceService) DeleteTeacher(ctx context.Context, req *pb.DeleteTeacherRequest) (*pb.DeleteTeacherReply, error) {
	if req.Id == 0 {
		return nil, errors.New("missing teacher ID forfor delete")
	}
	e := s.teacherUsecase.Delete(ctx, req.Id)
	if e != nil {
		return nil, e
	}
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

	go func() {
		if data, e := json.Marshal(reply); e == nil {
			_ = s.data.RedisCache.Set(ctx, key, data, 10*time.Minute)
		}
	}()
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
