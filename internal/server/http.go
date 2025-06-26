package server

import (
	v1 "DemoApp/api/helloworld/v1"
	"DemoApp/internal/conf"
	"DemoApp/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, class *service.ClassServiceService, student *service.StudentServiceService, teacher *service.TeacherServiceService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterStudentServiceHTTPServer(srv, student)
	v1.RegisterClassServiceHTTPServer(srv, class)
	v1.RegisterTeacherServiceHTTPServer(srv, teacher)
	return srv
}
