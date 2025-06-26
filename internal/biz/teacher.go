package biz

import (
	pb "DemoApp/api/helloworld/v1"
	common "DemoApp/internal/common"
	"DemoApp/internal/common/filter"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Teacher struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int32  `json:"age"`
	Class_id   int64  `json:"class_id"`
	Class_name string `json:"class_name"`
}

type TeacherRepo interface {
	FindAll(ctx context.Context, filter filter.TeacherFilter, pagination common.Pagination) ([]*Teacher, error)
	FindByID(ctx context.Context, id int64) (*Teacher, error)
	Create(ctx context.Context, input *Teacher) (*Teacher, error)
	Update(ctx context.Context, input *Teacher, id int64) error
	Delete(ctx context.Context, id int64) error
}

type TeacherUseCase struct {
	repo TeacherRepo
	log  *log.Helper
}

func (t *Teacher) To() *pb.TeacherReply {
	return &pb.TeacherReply{
		Id:        t.Id,
		Name:      t.Name,
		Email:     t.Email,
		Age:       t.Age,
		ClassName: t.Class_name,
	}
}

func NewTeacherUsecase(repo TeacherRepo, logger log.Logger) *TeacherUseCase {
	return &TeacherUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *TeacherUseCase) FindAll(ctx context.Context, filter filter.TeacherFilter, pagination common.Pagination) ([]*Teacher, error) {
	return uc.repo.FindAll(ctx, filter, pagination)
}
