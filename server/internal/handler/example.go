// Package handler contains GRPC handlers.

package handler

import (
	"context"
	"template-api-go/example"
	"template-api-go/example/pb/fakeapi"
)

type ExampleServer struct {
	DB example.DataFetcher
	PS example.DataNotifier
	fakeapi.UnimplementedFakeServiceServer
}

func (e ExampleServer) GetFakeRequest(ctx context.Context, request *fakeapi.FakeRequest) (*fakeapi.FakeResponse, error) {
	exampleData, err := e.DB.GetAllExampleData(ctx)
	if err != nil {
		return nil, err
	}

	err = e.PS.NotifyExampleData(ctx, exampleData[0])
	if err != nil {
		return nil, err
	}

	return &fakeapi.FakeResponse{}, nil
}
