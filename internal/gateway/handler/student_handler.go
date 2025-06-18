// internal/gateway/handler/student_handler.go
package handler

import (
	v1 "DemoApp/api/helloworld/v1"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// StudentHandler xử lý các yêu cầu liên quan đến student
type StudentHandler struct {
	client v1.StudentServiceClient
	v1.UnimplementedStudentServiceServer
}

// NewStudentHandler tạo một StudentHandler mới với gRPC client
func NewStudentHandler(endpoint string) (*StudentHandler, error) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	client := v1.NewStudentServiceClient(conn)
	return &StudentHandler{client: client}, nil
}

// CreateStudent forward yêu cầu tạo student đến service chính
func (h *StudentHandler) CreateStudent(ctx context.Context, req *v1.CreateStudentRequest) (*v1.CreateStudentReply, error) {
	return h.client.CreateStudent(ctx, req)
}

// GetStudent forward yêu cầu lấy thông tin student
func (h *StudentHandler) GetStudent(ctx context.Context, req *v1.GetStudentRequest) (*v1.GetStudentReply, error) {
	return h.client.GetStudent(ctx, req)
}

func (h *StudentHandler) ListStudent(ctx context.Context, req *v1.ListStudentRequest) (*v1.ListStudentReply, error) {
	return h.client.ListStudent(ctx, req)
}

func (h *StudentHandler) UpdateStudent(ctx context.Context, req *v1.UpdateStudentRequest) (*v1.UpdateStudentReply, error) {
	return h.client.UpdateStudent(ctx, req)
}

func (h *StudentHandler) DeleteStudent(ctx context.Context, req *v1.DeleteStudentRequest) (*v1.DeleteStudentReply, error) {
	return h.client.DeleteStudent(ctx, req)
}
