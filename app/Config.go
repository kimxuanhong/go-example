package app

import (
	"github.com/kimxuanhong/go-http/server"
	"github.com/kimxuanhong/go-postgres/postgres"
	"github.com/kimxuanhong/go-redis/redis"
)

type Config struct {
	Server   *server.Config   `yaml:"server"`
	Redis    *redis.Config    `yaml:"redis,omitempty"`
	Postgres *postgres.Config `yaml:"postgres,omitempty"`
}
