package domain

// ConsumerInfo represents consumer information from consumer service
type ConsumerInfo struct {
	ID       string
	IsActive bool
	Type     string
	Metadata map[string]interface{}
}
