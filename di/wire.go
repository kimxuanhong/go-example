//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kimxuanhong/go-example/internal/infrastructure/repository"
	"github.com/kimxuanhong/go-example/internal/interface/handler"
	"github.com/kimxuanhong/go-example/internal/usecase"
	"github.com/kimxuanhong/go-http/server"
	"github.com/kimxuanhong/go-postgres/postgres"
	"github.com/kimxuanhong/go-redis/redis"
	"github.com/kimxuanhong/go-utils/config"
)

type Config struct {
	Server   *server.Config   `yaml:"server"`
	Redis    *redis.Config    `yaml:"redis,omitempty"`
	Postgres *postgres.Config `yaml:"postgres,omitempty"`
}

type App struct {
	Cfg         *Config
	Server      server.Server
	UserHandler *handler.UserHandler
}

func InitApp() (*App, error) {
	wire.Build(
		LoadConfig,
		InitHttpServer,
		InitPostgres,
		repository.NewUserRepo,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}

// Tạo một provider cho config
func LoadConfig() (*Config, error) {
	cfg, err := config.LoadConfig[Config]()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Server luôn được khởi tạo từ cấu hình
func InitHttpServer(cfg *Config) (server.Server, error) {
	return server.NewServer(cfg.Server), nil
}

// Postgres chỉ được khởi tạo nếu có config postgres
func InitPostgres(cfg *Config) (postgres.Postgres, error) {
	return postgres.NewPostgres(cfg.Postgres)
}

// Redis chỉ được khởi tạo nếu có config redis
func InitRedis(cfg *Config) (redis.Redis, error) {
	return redis.NewRedis(cfg.Redis)
}
