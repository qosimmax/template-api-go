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
		Header:  header(headers),
		Data:    data,
	}

	_, err := c.PublishMsg(msg)
	if err != nil {
		return err
	}

	return nil
}

func header(h propagation.HeaderCarrier) nats.Header {
	if h == nil {
		return nil
	}

	// Find total number of values.
	nv := 0
	for _, vv := range h {
		nv += len(vv)
	}

	sv := make([]string, nv) // shared backing array for headers' values
	h2 := make(nats.Header, len(h))

	for k, vv := range h {
		if vv == nil {
			// Preserve nil values. ReverseProxy distinguishes
			// between nil and zero-length header values.
			h2[k] = nil
			continue
		}

		n := copy(sv, vv)
		h2[k] = sv[:n:n]
		sv = sv[n:]
	}

	return h2
}
