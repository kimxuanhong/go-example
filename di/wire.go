//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	dbCore "github.com/kimxuanhong/go-database/core"
	db "github.com/kimxuanhong/go-database/database"
	"github.com/kimxuanhong/go-example/internal/domain/validator"
	"github.com/kimxuanhong/go-example/internal/facade"
	"github.com/kimxuanhong/go-example/internal/infrastructure/external"
	"github.com/kimxuanhong/go-example/internal/infrastructure/repository"
	"github.com/kimxuanhong/go-example/internal/interface/handler"
	"github.com/kimxuanhong/go-example/internal/usecase"
	"github.com/kimxuanhong/go-example/pkg"
	"github.com/kimxuanhong/go-server/core"
	"github.com/kimxuanhong/go-server/jwt"
	"github.com/kimxuanhong/go-server/server"
	"github.com/kimxuanhong/go-utils/config"
)

type Config struct {
	Server          *core.Config   `mapstructure:"server"`
	Postgres        *dbCore.Config `mapstructure:"postgres,omitempty"`
	ReplicaPostgres *dbCore.Config `mapstructure:"replica_postgres,omitempty"`
	Jwt             *jwt.Config    `mapstructure:"jwt,omitempty"`
}

// ConfigSet chứa các provider liên quan đến cấu hình
var ConfigSet = wire.NewSet(
	LoadConfig,
	InitHttpServer,
)

// DatabaseSet chứa các provider liên quan đến database
var DatabaseSet = wire.NewSet(
	InitPostgres,
	InitReplicaPostgres,
)

// RepositorySet chứa các provider liên quan đến repositories
var RepositorySet = wire.NewSet(
	repository.NewUserRepo,
)

// UsecaseSet chứa các provider liên quan đến usecases và facades
var UsecaseSet = wire.NewSet(
	validator.NewUserValidator,
	usecase.NewUserUsecase,
	facade.NewUserFacade,
)

// HandlerSet chứa các provider liên quan đến handlers
var HandlerSet = wire.NewSet(
	handler.NewUserHandler,
)

type Handlers []interface{}

type App struct {
	Cfg      *Config
	Server   core.Server
	Handlers Handlers
}

func ProvideHandlers(user *handler.UserHandler) Handlers {
	return Handlers{user}
}

func InitApp() (*App, error) {
	wire.Build(
		ConfigSet,
		DatabaseSet,
		RepositorySet,
		UsecaseSet,
		external.NewAccountClient,
		HandlerSet,
		ProvideHandlers,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}

// LoadConfig tạo một provider cho config
func LoadConfig() (*Config, error) {
	cfg, err := config.LoadConfig[Config]()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// InitHttpServer khởi tạo HTTP server từ cấu hình
func InitHttpServer(cfg *Config) (core.Server, error) {
	return server.NewServer(cfg.Server), nil
}

// InitPostgres khởi tạo Postgres nếu có config postgres
func InitPostgres(cfg *Config) (pkg.MainPostgres, error) {
	return db.NewDatabase(cfg.Postgres)
}

// InitReplicaPostgres khởi tạo Replica Postgres nếu có config
func InitReplicaPostgres(cfg *Config) (pkg.ReplicaPostgres, error) {
	return db.NewDatabase(cfg.ReplicaPostgres)
}
