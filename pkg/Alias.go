package pkg

import (
	database "github.com/kimxuanhong/go-database/core"
	"github.com/kimxuanhong/go-http/client"
)

type MainPostgres database.Database
type ReplicaPostgres database.Database
type AccountClient client.Client
type ConsumerClient client.Client
