package pubsub

import (
	ps "cloud.google.com/go/pubsub"
	"context"
	"fmt"

	"template-api-go/config"
)

// Client holds the PubSub client.
type Client struct {
	PubSubClient *ps.Client
}

// Init sets up a new pubsub client.
func (c *Client) Init(config *config.Config) error {
	psClient, err := ps.NewClient(context.Background(), config.PubSubProjectName)
	if err != nil {
		return fmt.Errorf("error creating pubsub client: %w", err)
	}
	c.PubSubClient = psClient
	return nil
}

func (c *Client) send(ctx context.Context, topicName string, data []byte) error {
	return nil
}
