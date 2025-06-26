package biz

import (
	common "DemoApp/internal/common"
	"context"
)

type IRepository[T any] interface {
	Create(ctx context.Context, input *T) (*T, error)
	Update(ctx context.Context, input *T, id int64) error
	FindAll(ctx context.Context, filter any, pagination common.Pagination) ([]*T, error)
	Delete(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (*T, error)
}
