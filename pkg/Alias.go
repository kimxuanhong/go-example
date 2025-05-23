package pkg

import (
	"github.com/kimxuanhong/go-database/db"
	"github.com/kimxuanhong/go-http/client"
)

type MainPostgres *db.Database
type ReplicaPostgres *db.Database
type AccountClient client.Client
type ConsumerClient client.Client
