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
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       int32  `json:"age"`
	ClassID   int64  `json:"class_id"`
	ClassName string `json:"class_name"`
}

// UpdateTeacher represents the fields that can be optionally updated for a teacher.
// Only non-nil fields will be applied during update.
type UpdateTeacher struct {
	Name    *string `json:"id"`
	Email   *string `json:"email"`
	Age     *int32  `json:"name"`
	ClassID *int64  `json:"class_id"`
}

// TeacherRepo defines the methods required for accessing teacher data from the data layer.
type TeacherRepo interface {
	FindAll(ctx context.Context, filter filter.TeacherFilter, pagination common.Pagination) ([]*Teacher, error)
	FindByID(ctx context.Context, id int64) (*Teacher, error)
	Create(ctx context.Context, input *Teacher) (*Teacher, error)
	Update(ctx context.Context, input *UpdateTeacher, id int64) error
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

// FindByID retrieves a single teacher that match the provided id
func (uc *TeacherUseCase) FindByID(ctx context.Context, id int64) (*Teacher, error) {
	return uc.repo.FindByID(ctx, id)
}

// Create a teacher with data from input
// return a new teacher entity
func (uc *TeacherUseCase) Create(ctx context.Context, input *Teacher) (*Teacher, error) {
	return uc.repo.Create(ctx, input)
}

// Update a existed teacher with data from input that match the provided id
// return a teacher entity with new data
func (uc *TeacherUseCase) Update(ctx context.Context, input *UpdateTeacher, id int64) error {
	return uc.repo.Update(ctx, input, id)
}

// Delete a existed teacher from id input return error
func (uc *TeacherUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
