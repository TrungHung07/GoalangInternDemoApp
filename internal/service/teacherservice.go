package service

import (
	"context"
	// "fmt"

	pb "DemoApp/api/helloworld/v1"
	data "DemoApp/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

type TeacherServiceService struct {
	pb.UnimplementedTeacherServiceServer
	data *data.Data
	log  *log.Helper
}

func NewTeacherServiceService(data *data.Data, logger log.Logger) *TeacherServiceService {
	return &TeacherServiceService{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *TeacherServiceService) GetTeacher(ctx context.Context, req *pb.GetTeacherRequest) (*pb.TeacherReply, error) {

	return &pb.TeacherReply{}, nil
}
func (s *TeacherServiceService) CreateTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*pb.CreateTeacherReply, error) {
	teacher, err := s.data.DB.Teacher.Create().
		SetName(req.Name).
		SetEmail(req.Email).
		SetClassID(int(req.ClassId)).
		Save(ctx)
	if err != nil {
		s.log.Error("Failed to create teacher:", err)
		return nil, err
	}
	s.log.Infof("Created teacher: %v", teacher)
	return &pb.CreateTeacherReply{Message: "Thêm thành công !"}, nil
}

func (s *TeacherServiceService) UpdateTeacher(ctx context.Context, req *pb.UpdateTeacherRequest) (*pb.UpdateTeacherReply, error) {

	teacher, err := s.data.DB.Teacher.Get(ctx, int(req.Id))
	s.log.Infof("teacher and error :" + err.Error())
	if err != nil {
		s.log.Error("\nFailed to find teacher ( if is correct !):\n", err)
		return &pb.UpdateTeacherReply{Message: "Không tìm thấy id!"}, nil
	}
	update := teacher.Update()
	if req.Name != "" {
		update.SetName(req.Name)
	}
	if req.Email != "" {
		update.SetEmail(req.Email)
	}
	teacherUpdateed, err := update.Save(ctx)

	if err != nil {
		return &pb.UpdateTeacherReply{Message: "Cập nhật thất bại!"}, err
	}
	// teacher, err := s.data.DB.Teacher.UpdateOneID(int(req.Id)).
	// 	SetName(req.Name).
	// 	SetEmail(req.Email).
	// 	SetGrade(int(req.Grade)).
	// 	Save(ctx)
	s.log.Infof("Updated teacher: %v", teacherUpdateed)
	return &pb.UpdateTeacherReply{Message: "Cập nhật thành công rồi đó !"}, nil
}

func (s *TeacherServiceService) DeleteTeacher(ctx context.Context, req *pb.DeleteTeacherRequest) (*pb.DeleteTeacherReply, error) {

	err := s.data.DB.Teacher.DeleteOneID(int(req.Id)).Exec(ctx)
	if err != nil {
		s.log.Error("Failed to delete teacher:", err)
		return nil, err
	}
	s.log.Infof("Deleted teacher with ID: %d", req.Id)

	return &pb.DeleteTeacherReply{Message: "Xoá thành công !"}, nil
}

func (s *TeacherServiceService) ListTeachers(ctx context.Context, req *pb.ListTeachersRequest) (*pb.ListTeachersReply, error) {
	teachers, err := s.data.DB.Teacher.Query().All(ctx)
	// fmt.Println("Teachers in DB:", teachers)
	if err != nil {
		s.log.Error("Failed to list teachers:", err)
		return nil, err
	}
	var teacherReplies []*pb.TeacherReply
	for _, teacher := range teachers {
		teacherReplies = append(teacherReplies, &pb.TeacherReply{
			Id:        int64(teacher.ID),
			Name:      teacher.Name,
			Email:     teacher.Email,
			ClassName: teacher.Edges.Classes.Name,
		})
	}
	return &pb.ListTeachersReply{Teachers: teacherReplies}, nil
}
