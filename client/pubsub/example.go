package pubsub

import (
	"context"
	"fmt"
	"template-api-go/example"
	"template-api-go/example/pb/fakeapi"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/protobuf/proto"

	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("pubsub")

// NotifyExampleData is used to publish example data to pubsub.
func (c *Client) NotifyExampleData(ctx context.Context, exampleData example.Data) error {
	ctx, span := tracer.Start(ctx, "NotifyExampleData")
	defer span.End()

	fakeData := &fakeapi.FakeData{
		IsFake: exampleData.IsFake,
		Date:   timestamppb.New(exampleData.Date),
	}

	data, err := proto.Marshal(fakeData)
	if err != nil {
		return fmt.Errorf("error marshalling example data to send to pubsub: %w", err)
	}
	err = c.send(ctx, "example", data)
	if err != nil {
		return fmt.Errorf("error sending example data message to pubsub: %w", err)
	}
	return nil
}
