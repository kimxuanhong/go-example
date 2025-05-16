package external

import "github.com/kimxuanhong/go-feign/feign"

type AccountClient struct {
	_       struct{}                                           `feign:"@Url http://localhost:8081/api/v1"`
	GetUser func(id string, username string) (*UserRes, error) `feign:"@GET /users/{id} | @Path id | @Query username"`
}

func NewAccountClient() *AccountClient {
	client := feign.NewClient()
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
