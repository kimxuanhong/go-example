package pkg

import (
	"github.com/kimxuanhong/go-http/client"
	"github.com/kimxuanhong/go-postgres/postgres"
)

type MainPostgres postgres.Postgres
type ReplicaPostgres postgres.Postgres
type AccountClient client.Client
type ConsumerClient client.Client
