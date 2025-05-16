package external

import (
	"context"
	"github.com/kimxuanhong/go-example/internal/domain"
)

type FeignConsumerClient struct {
	GetConsumer func(ctx context.Context, username, include_metadata, source string) (*consumerInfoResponse, error) `feign:"@GET /consumers/{username} | @Path username | @Query include_metadata | @Query source"`
}

type ConsumerClient struct {
	client *FeignConsumerClient
}

func NewConsumerClient(client *FeignConsumerClient) *ConsumerClient {
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
	resp, err := c.client.GetConsumer(ctx, userName, "true", opts.Source)
	if err != nil {
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
