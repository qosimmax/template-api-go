package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"template-api-go/example"

	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("pubsub")

// NotifyExampleData is used to publish example data to pubsub.
func (c *Client) NotifyExampleData(ctx context.Context, exampleData example.Data) error {
	ctx, span := tracer.Start(ctx, "NotifyExampleData")
	defer span.End()

	data, err := json.Marshal(exampleData)
	if err != nil {
		return fmt.Errorf("error marshalling example data to send to pubsub: %w", err)
	}
	err = c.send(ctx, "example", data)
	if err != nil {
		return fmt.Errorf("error sending example data message to pubsub: %w", err)
	}
	return nil
}
