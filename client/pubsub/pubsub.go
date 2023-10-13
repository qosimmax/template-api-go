package pubsub

import (
	"context"
	"template-api-go/config"
)

// Client holds the PubSub client.
type Client struct {
}

// Init sets up a new pubsub client.
func (c *Client) Init(config *config.Config) error {
	return nil
}

func (c *Client) send(ctx context.Context, topicName string, data []byte) error {
	return nil
}
