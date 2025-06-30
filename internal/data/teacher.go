package data

import (
	"DemoApp/ent"
	"DemoApp/ent/teacher"
	"DemoApp/internal/biz"
	"DemoApp/internal/common"
	"DemoApp/internal/common/filter"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.TeacherRepo = (*teacherRepo)(nil)

type teacherRepo struct {
	data *Data
	log  *log.Helper
}

// NewTeacherRepo creates and returns a new instance of TeacherRepo using the provided data layer and logger.
func NewTeacherRepo(data *Data, logger log.Logger) biz.TeacherRepo {
	return &teacherRepo{
		data: data,
		log:  log.NewHelper(logger),	
	}
} 

func applyPagination(query *ent.TeacherQuery, p common.Pagination) *ent.TeacherQuery {
	return query.Limit(p.Limit()).Offset(p.Offset())
}

func applyTeacherFilter(query *ent.TeacherQuery, f filter.TeacherFilter) *ent.TeacherQuery {
	if f.Email != nil {
		query = query.Where(teacher.EmailContains(*f.Email))
	}
	if f.MinAge != nil {
		query = query.Where(teacher.AgeGTE(int(*f.MinAge)))
	}
	if f.MaxAge != nil {
		query = query.Where(teacher.AgeLTE(int(*f.MaxAge)))
	}
	return query
}

func (r *teacherRepo) FindAll(ctx context.Context, filter filter.TeacherFilter, pagination common.Pagination) ([]*biz.Teacher, error) {
	// var query *ent.TeacherQuery
	query := r.data.DB.Teacher.Query()
	query = applyTeacherFilter(query, filter)
	query = applyPagination(query, pagination)

	entities, err := query.WithClasses().All(ctx)

	if err != nil {
		return nil, err
	}

	//Map sang biz.Teacher
	var result []*biz.Teacher
	for _, t := range entities {
		var className string
		if t.Edges.Classes != nil {
			className = t.Edges.Classes.Name
		}

		result = append(result, &biz.Teacher{
			Id:         int64(t.ID),
			Name:       t.Name,
			Email:      t.Email,
			Age:        int32(t.Age),
			Class_id:   int64(t.ClassID),
			Class_name: className, // <-- bổ sung field nếu bạn có trong `biz.Teacher`
		})
	}
	return result, nil
}

func (r *teacherRepo) FindByID(ctx context.Context, id int64) (*biz.Teacher, error) {
	// TODO: implement
	return nil, nil
}

func (r *teacherRepo) Create(ctx context.Context, input *biz.Teacher) (*biz.Teacher, error) {
	// TODO: implement
	return nil, nil
}

func (r *teacherRepo) Update(ctx context.Context, input *biz.Teacher, id int64) error {
	// TODO: implement
	return nil
}

func (r *teacherRepo) Delete(ctx context.Context, id int64) error {
	// TODO: implement
	return nil
}
