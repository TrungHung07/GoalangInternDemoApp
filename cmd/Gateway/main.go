package main

import (
	"context"
	"flag"
	"os"

	v1 "DemoApp/api/helloworld/v1"
	"DemoApp/internal/conf"
	gateway_config "DemoApp/internal/gateway/config"
	"DemoApp/internal/gateway/handler"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/joho/godotenv"
)

var (
	Name     = "gateway"
	Version  = "v0.0.1"
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs/gateway_config.yaml", "config path, eg: -conf gateway_config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	if err := godotenv.Load("../.env"); err != nil { // Đường dẫn đến file .env
		log.NewStdLogger(os.Stdout).Log(log.LevelWarn, "msg", "Cannot load .env file, proceeding without it")
	}

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	c := config.New(
		config.WithSource(file.NewSource(flagconf)),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// Tạo cấu hình gateway
	gwConfig, err := gateway_config.NewConfig(c)
	if err != nil {
		panic(err)
	}

	log.Info("\n Gateway Config", "\n HTTPAddr", gwConfig.Server.HTTP.Addr, "\n GRPCAddr", gwConfig.Server.GRPC.Addr, "\n ServiceEndpoint", gwConfig.Data.ServiceEndpoint)
	if gwConfig.Data.ServiceEndpoint == "" {
		log.Error("ServiceEndpoint is empty, check config.yaml")
		panic("ServiceEndpoint is not configured")
	}

	// Khởi tạo các handler
	classHandler, err := handler.NewClassHandler(gwConfig.Data.ServiceEndpoint)
	if err != nil {
		panic(err)
	}
	studentHandler, err := handler.NewStudentHandler(gwConfig.Data.ServiceEndpoint)
	if err != nil {
		panic(err)
	}

	// Khởi tạo HTTP server
	httpSrv := http.NewServer(
		http.Address(gwConfig.Server.HTTP.Addr),
	)
	// if err := httpSrv.Start(context.Background()); err != nil {
	// 	panic(err)
	// }

	// Khởi tạo gRPC server
	grpcSrv := grpc.NewServer(
		grpc.Address(gwConfig.Server.GRPC.Addr),
	)

	// if err := grpcSrv.Start(context.Background()); err != nil {
	// 	panic(err)
	// }

	// Đăng ký các service gRPC
	v1.RegisterClassServiceServer(grpcSrv, classHandler)
	v1.RegisterStudentServiceServer(grpcSrv, studentHandler)
	v1.RegisterClassServiceHTTPServer(httpSrv, classHandler)
	v1.RegisterStudentServiceHTTPServer(httpSrv, studentHandler)

	// Tạo ứng dụng Kratos
	app := newApp(logger, grpcSrv, httpSrv)

	// Hàm dọn dẹp
	cleanup := func() {
		log.Info("Closing the gateway resources")
		_ = httpSrv.Stop(context.Background())
		_ = grpcSrv.Stop(context.Background())
	}
	defer cleanup()

	// Run ứng dụng
	if err := app.Run(); err != nil {
		panic(err)
	}
}
