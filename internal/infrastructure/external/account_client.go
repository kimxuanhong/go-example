package external

import (
	"context"
	"github.com/kimxuanhong/go-feign/feign"
)

type AccountClient struct {
	GetUser func(ctx context.Context, id string, username string) (*UserRes, error) `feign:"@GET /users/{id} | @Path id | @Query username"`
}

func NewAccountClient(config *feign.Config) *AccountClient {
	client := feign.NewClient(config)
	accountClient := &AccountClient{}
	client.Create(accountClient)
	return accountClient
}

type UserRes struct {
	ID        string
	PartnerId string
	Total     int
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Status    string
}
