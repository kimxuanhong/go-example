package external

import (
	"context"
	"fmt"

	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/pkg"
)

type AccountClient struct {
	client pkg.AccountClient
}

func NewAccountClient(client pkg.AccountClient) *AccountClient {
	return &AccountClient{
		client: client,
	}
}

type GetUserOptions struct {
	IncludeProfile bool
	Fields         []string
}

type accountUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Profile  *struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"profile,omitempty"`
}

// GetUser gọi API lấy thông tin user từ account service
func (c *AccountClient) GetUser(ctx context.Context, userName string, opts *GetUserOptions) (*domain.User, error) {
	// Xây dựng path với query params
	path := fmt.Sprintf("/users/%s", userName)
	if opts != nil {
		if opts.IncludeProfile {
			path += "?include_profile=true"
		}
		if len(opts.Fields) > 0 {
			path += fmt.Sprintf("&fields=[%s]", opts.Fields)
		}
	}

	// Gọi API thông qua base client
	var resp accountUserResponse
	if err := c.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	// Transform response data sang domain model
	user := &domain.User{
		UserName: resp.Username,
		Email:    resp.Email,
		Status:   resp.Status,
	}

	// Parse thêm profile data nếu có
	if opts != nil && opts.IncludeProfile && resp.Profile != nil {
		user.FirstName = resp.Profile.FirstName
		user.LastName = resp.Profile.LastName
	}

	return user, nil
}
