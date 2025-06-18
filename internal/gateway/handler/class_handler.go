// internal/gateway/handler/class_handler.go
package handler

import (
	v1 "DemoApp/api/helloworld/v1" // Đảm bảo import đúng package proto
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// ClassHandler xử lý các yêu cầu liên quan đến class
type ClassHandler struct {
	client v1.ClassServiceClient
	v1.UnimplementedClassServiceServer
}

// NewClassHandler tạo một ClassHandler mới với gRPC client
func NewClassHandler(endpoint string) (*ClassHandler, error) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	client := v1.NewClassServiceClient(conn)
	return &ClassHandler{client: client}, nil
}

// CreateClass forward yêu cầu tạo class đến service chính
func (h *ClassHandler) CreateClass(ctx context.Context, req *v1.CreateClassRequest) (*v1.CreateClassReply, error) {
	return h.client.CreateClass(ctx, req)
}

// GetClass forward yêu cầu lấy thông tin class
func (h *ClassHandler) GetClass(ctx context.Context, req *v1.GetClassRequest) (*v1.GetClassReply, error) {
	return h.client.GetClass(ctx, req)
}

func (h *ClassHandler) ListClass(ctx context.Context, req *v1.ListClassRequest) (*v1.ListClassReply, error) {
	return h.client.ListClass(ctx, req)
}

func (h *ClassHandler) UpdateClass(ctx context.Context, req *v1.UpdateClassRequest) (*v1.UpdateClassReply, error) {
	return h.client.UpdateClass(ctx, req)
}

func (h *ClassHandler) DeleteClass(ctx context.Context, req *v1.DeleteClassRequest) (*v1.DeleteClassReply, error) {
	return h.client.DeleteClass(ctx, req)
}
