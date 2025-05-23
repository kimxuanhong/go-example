//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kimxuanhong/go-database/db"
	"github.com/kimxuanhong/go-example/internal/delivery/http"
	"github.com/kimxuanhong/go-example/internal/domain/rules"
	"github.com/kimxuanhong/go-example/internal/facade"
	"github.com/kimxuanhong/go-example/internal/infrastructure/external"
	"github.com/kimxuanhong/go-example/internal/infrastructure/repository"
	"github.com/kimxuanhong/go-example/internal/usecase"
	"github.com/kimxuanhong/go-example/pkg"
	"github.com/kimxuanhong/go-feign/feign"
	"github.com/kimxuanhong/go-server/core"
	"github.com/kimxuanhong/go-server/jwt"
	"github.com/kimxuanhong/go-server/server"
	"github.com/kimxuanhong/go-utils/config"
)

type Config struct {
	Server              *core.Config  `mapstructure:"server"`
	Postgres            *db.Config    `mapstructure:"postgres,omitempty"`
	ReplicaPostgres     *db.Config    `mapstructure:"replica_postgres,omitempty"`
	Jwt                 *jwt.Config   `mapstructure:"jwt,omitempty"`
	AccountClientConfig *feign.Config `mapstructure:"account_client,omitempty"`
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
	rules.NewUserValidator,
	usecase.NewUserUsecase,
	facade.NewUserFacade,
)

// HandlerSet chứa các provider liên quan đến handlers
var HandlerSet = wire.NewSet(
	http.NewUserHandler,
)

type Handlers []interface{}

type App struct {
	Cfg      *Config
	Server   core.Server
	Handlers Handlers
}

func ProvideHandlers(user *http.UserHandler) Handlers {
	return Handlers{user}
}

func InitApp() (*App, error) {
	wire.Build(
		ConfigSet,
		DatabaseSet,
		RepositorySet,
		UsecaseSet,
		InitAccountClient,
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
	return db.Open(cfg.Postgres)
}

// InitReplicaPostgres khởi tạo Replica Postgres nếu có config
func InitReplicaPostgres(cfg *Config) (pkg.ReplicaPostgres, error) {
	return db.Open(cfg.ReplicaPostgres)
}

// InitReplicaPostgres khởi tạo Replica Postgres nếu có config
func InitAccountClient(cfg *Config) (*external.AccountClient, error) {
	return external.NewAccountClient(cfg.AccountClientConfig), nil
}
