package biz

import (
	pb "DemoApp/api/helloworld/v1"
	common "DemoApp/internal/common"
	"DemoApp/internal/common/filter"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Teacher represents the business model of a teacher with related class information.
type Teacher struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int32  `json:"age"`
	ClassID   int64  `json:"class_id"`
	ClassName string `json:"class_name"`
}

// TeacherRepo defines the methods required for accessing teacher data from the data layer.
type TeacherRepo interface {
	FindAll(ctx context.Context, filter filter.TeacherFilter, pagination common.Pagination) ([]*Teacher, error)
	FindByID(ctx context.Context, id int64) (*Teacher, error)
	Create(ctx context.Context, input *Teacher) (*Teacher, error)
	Update(ctx context.Context, input *Teacher, id int64) error
	Delete(ctx context.Context, id int64) error
}

// TeacherUseCase handles the business logic related to teachers.
type TeacherUseCase struct {
	repo TeacherRepo
	log  *log.Helper
}

// To converts a Teacher domain model to a protobuf-compatible TeacherReply for API responses.
func (t *Teacher) To() *pb.TeacherReply {
	return &pb.TeacherReply{
		Id:        t.ID,
		Name:      t.Name,
		Email:     t.Email,
		Age:       t.Age,
		ClassName: t.ClassName,
	} 
}

// NewTeacherUsecase creates and returns a new instance of TeacherUseCase with the given repository and logger.
func NewTeacherUsecase(repo TeacherRepo, logger log.Logger) *TeacherUseCase {
	return &TeacherUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// FindAll retrieves a list of teachers that match the provided filter and pagination options.
func (uc *TeacherUseCase) FindAll(ctx context.Context, filter filter.TeacherFilter, pagination common.Pagination) ([]*Teacher, error) {
	return uc.repo.FindAll(ctx, filter, pagination)
}
