//go:build !wireinject

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
package app

import (
	"github.com/kimxuanhong/go-http/server"
	"github.com/kimxuanhong/go-postgres/postgres"
	"github.com/kimxuanhong/go-redis/redis"
	"github.com/kimxuanhong/go-utils/config"
)

func InitApp() (*App, error) {
	cfg, err := config.LoadConfig[Config]()
	if err != nil {
		return nil, err
	}
	sev, err := provideServer(cfg)
	if err != nil {
		return nil, err
	}
	ptg, err := providePostgres(cfg)
	if err != nil {
		return nil, err
	}
	rds, err := provideRedis(cfg)
	if err != nil {
		return nil, err
	}
	app := &App{
		Server:   sev,
		Postgres: ptg,
		Redis:    rds,
	}
	return app, nil
}

type App struct {
	Server   server.Server
	Postgres *postgres.Postgres
	Redis    *redis.Redis
}

// Server luôn được khởi tạo từ cấu hình
func provideServer(cfg *Config) (server.Server, error) {
	return server.NewServer(cfg.Server), nil
}

// Postgres chỉ được khởi tạo nếu có config postgres
func providePostgres(cfg *Config) (*postgres.Postgres, error) {
	if cfg.Postgres == nil {
		return nil, nil
	}
	return postgres.NewPostgres(cfg.Postgres)
}

// Redis chỉ được khởi tạo nếu có config redis
func provideRedis(cfg *Config) (*redis.Redis, error) {
	if cfg.Redis == nil {
		return nil, nil
	}
	return redis.NewRedis(cfg.Redis)
}
