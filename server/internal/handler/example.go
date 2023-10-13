// Package handler contains GRPC handlers.

package handler

import (
	"context"
	"template-api-go/example"
	"template-api-go/example/pb/fakeapi"
)

type ExampleServer struct {
	DB example.DataFetcher
	fakeapi.UnimplementedFakeServiceServer
}

func (e ExampleServer) GetFakeRequest(ctx context.Context, request *fakeapi.FakeRequest) (*fakeapi.FakeResponse, error) {
	_, err := e.DB.GetAllExampleData(ctx)
	if err != nil {
		return nil, err
	}

	return &fakeapi.FakeResponse{}, nil
}
