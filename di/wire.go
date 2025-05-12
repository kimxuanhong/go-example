//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kimxuanhong/go-example/internal/domain/validator"
	"github.com/kimxuanhong/go-example/internal/facade"
	"github.com/kimxuanhong/go-example/internal/infrastructure/external"
	"github.com/kimxuanhong/go-example/internal/infrastructure/repository"
	"github.com/kimxuanhong/go-example/internal/interface/handler"
	"github.com/kimxuanhong/go-example/internal/usecase"
	"github.com/kimxuanhong/go-example/pkg"
	"github.com/kimxuanhong/go-http/client"
	"github.com/kimxuanhong/go-postgres/postgres"
	"github.com/kimxuanhong/go-redis/redis"
	"github.com/kimxuanhong/go-server/core"
	"github.com/kimxuanhong/go-server/gin"
	"github.com/kimxuanhong/go-utils/config"
)

type Config struct {
	Server          *gin.Config      `yaml:"server"`
	Redis           *redis.Config    `yaml:"redis,omitempty"`
	Postgres        *postgres.Config `yaml:"postgres,omitempty"`
	ReplicaPostgres *postgres.Config `yaml:"replica_postgres,omitempty"`
	AccountClient   *client.Config   `yaml:"account_client,omitempty"`
	ConsumerClient  *client.Config   `yaml:"consumer_client,omitempty"`
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
	InitRedis,
)

// ClientSet chứa các provider liên quan đến HTTP clients
var ClientSet = wire.NewSet(
	InitAccountClient,
	InitConsumerClient,
	external.NewAccountClient,
	external.NewConsumerClient,
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
		ClientSet,
		RepositorySet,
		UsecaseSet,
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
	return gin.NewServer(cfg.Server), nil
}

// InitPostgres khởi tạo Postgres nếu có config postgres
func InitPostgres(cfg *Config) (pkg.MainPostgres, error) {
	return postgres.NewPostgres(cfg.Postgres)
}

// InitReplicaPostgres khởi tạo Replica Postgres nếu có config
func InitReplicaPostgres(cfg *Config) (pkg.ReplicaPostgres, error) {
	return postgres.NewPostgres(cfg.ReplicaPostgres)
}

// InitRedis khởi tạo Redis nếu có config redis
func InitRedis(cfg *Config) (redis.Redis, error) {
	return redis.NewRedis(cfg.Redis)
}

// InitAccountClient khởi tạo Account HTTP client
func InitAccountClient(cfg *Config) pkg.AccountClient {
	return client.NewClient(cfg.AccountClient)
}

// InitConsumerClient khởi tạo Consumer HTTP client
func InitConsumerClient(cfg *Config) pkg.ConsumerClient {
	return client.NewClient(cfg.ConsumerClient)
}
