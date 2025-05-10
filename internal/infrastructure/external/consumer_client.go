package external

import (
	"context"
	"fmt"
	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/pkg"
)

type ConsumerClient struct {
	client pkg.ConsumerClient
}

func NewConsumerClient(client pkg.ConsumerClient) *ConsumerClient {
	return &ConsumerClient{
		client: client,
	}
}

type GetConsumerInfoOptions struct {
	IncludeMetadata bool
	Source          string
}

type consumerInfoResponse struct {
	ID       string                 `json:"id"`
	IsActive bool                   `json:"is_active"`
	Type     string                 `json:"type"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// GetConsumerInfo lấy thông tin consumer từ consumer service
func (c *ConsumerClient) GetConsumerInfo(ctx context.Context, userName string, opts *GetConsumerInfoOptions) (*domain.ConsumerInfo, error) {
	// Xây dựng path với query params
	path := fmt.Sprintf("/consumers/%s", userName)
	if opts != nil {
		if opts.IncludeMetadata {
			path += "?include_metadata=true"
		}
		if opts.Source != "" {
			path += fmt.Sprintf("&source=%s", opts.Source)
		}
	}

	// Gọi API thông qua base client
	var resp consumerInfoResponse
	if err := c.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	// Transform response data sang domain model
	consumerInfo := &domain.ConsumerInfo{
		ID:       resp.ID,
		IsActive: resp.IsActive,
		Type:     resp.Type,
	}

	// Parse thêm metadata nếu có
	if opts != nil && opts.IncludeMetadata {
		consumerInfo.Metadata = resp.Metadata
	}

	return consumerInfo, nil
}
