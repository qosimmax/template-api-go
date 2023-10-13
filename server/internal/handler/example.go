// Package handler contains GRPC handlers.

package handler

import (
	"context"
	"template-api-go/example"
	"template-api-go/proto/fake-api"
)

type ExampleServer struct {
	DB example.DataFetcher
	fakeapi.UnimplementedExampleServer
}

func (e ExampleServer) GetExampleRequest(ctx context.Context, request *fakeapi.ExampleRequest) (*fakeapi.ExampleResponse, error) {
	_, err := e.DB.GetAllExampleData(ctx)
	if err != nil {
		return nil, err
	}

	return &fakeapi.ExampleResponse{}, nil
}
