package gateway_config

import "github.com/go-kratos/kratos/v2/config"

// Config đại diện cho cấu hình của gateway
type Config struct {
	Data struct {
		ServiceEndpoint string `yaml:"service_endpoint"`
		Database        struct {
			Driver string `yaml:"driver"`
			Source string `yaml:"source"`
		} `yaml:"database"`
		Redis struct {
			Addr         string `yaml:"addr"`
			ReadTimeout  string `yaml:"read_timeout"`
			WriteTimeout string `yaml:"write_timeout"`
		} `yaml:"redis"`
	} `yaml:"data"`
	Server struct {
		HTTP struct {
			Addr    string `yaml:"addr"`
			Timeout string `yaml:"timeout"`
		} `yaml:"http"`
		GRPC struct {
			Addr    string `yaml:"addr"`
			Timeout string `yaml:"timeout"`
		} `yaml:"grpc"`
	} `yaml:"server"`
}

// NewConfig tạo một cấu hình mới từ nguồn config của Kratos
func NewConfig(c config.Config) (*Config, error) {
	var cfg Config
	if err := c.Scan(&cfg); err != nil {
		return nil, err
	}
	// Đặt giá trị mặc định nếu các trường rỗng
	if cfg.Server.HTTP.Addr == "" {
		cfg.Server.HTTP.Addr = "0.0.0.0:8001"
	}
	if cfg.Server.GRPC.Addr == "" {
		cfg.Server.GRPC.Addr = "0.0.0.0:9001"
	}
	if cfg.Data.ServiceEndpoint == "" {
		cfg.Data.ServiceEndpoint = "localhost:9000"
	}
	if cfg.Data.Database.Source == "" {
		cfg.Data.Database.Source = "user=postgres password=123 dbname=test host=localhost port=5432 sslmode=disable" // Giá trị mặc định
	}
	return &cfg, nil
}
