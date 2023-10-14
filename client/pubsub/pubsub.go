package pubsub

import (
	"context"
	"template-api-go/config"
	"template-api-go/monitoring/trace"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/propagation"
)

// Client holds the PubSub client.
type Client struct {
	nats.JetStreamContext
}

// Init sets up a new pubsub client.
func (c *Client) Init(config *config.Config) error {
	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		return err
	}

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(10000))
	if err != nil {
		return err
	}

	c.JetStreamContext = js
	return nil
}

func (c *Client) send(ctx context.Context, topicName string, data []byte) error {
	headers := make(propagation.HeaderCarrier)
	trace.InjectIntoCarrier(ctx, headers)

	msg := &nats.Msg{
		Subject: topicName,
		Header:  nats.Header(headers),
		Data:    data,
	}

	_, err := c.PublishMsg(msg)
	if err != nil {
		return err
	}

	return nil
}
