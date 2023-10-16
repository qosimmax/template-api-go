// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: service.proto

package fakeapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FakeService_GetFakeRequest_FullMethodName = "/fakeapi.FakeService/GetFakeRequest"
)

// FakeServiceClient is the client API for FakeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FakeServiceClient interface {
	GetFakeRequest(ctx context.Context, in *FakeRequest, opts ...grpc.CallOption) (*FakeResponse, error)
}

type fakeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFakeServiceClient(cc grpc.ClientConnInterface) FakeServiceClient {
	return &fakeServiceClient{cc}
}

func (c *fakeServiceClient) GetFakeRequest(ctx context.Context, in *FakeRequest, opts ...grpc.CallOption) (*FakeResponse, error) {
	out := new(FakeResponse)
	err := c.cc.Invoke(ctx, FakeService_GetFakeRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FakeServiceServer is the server API for FakeService service.
// All implementations must embed UnimplementedFakeServiceServer
// for forward compatibility
type FakeServiceServer interface {
	GetFakeRequest(context.Context, *FakeRequest) (*FakeResponse, error)
	mustEmbedUnimplementedFakeServiceServer()
}

// UnimplementedFakeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFakeServiceServer struct {
}

func (UnimplementedFakeServiceServer) GetFakeRequest(context.Context, *FakeRequest) (*FakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFakeRequest not implemented")
}
func (UnimplementedFakeServiceServer) mustEmbedUnimplementedFakeServiceServer() {}

// UnsafeFakeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FakeServiceServer will
// result in compilation errors.
type UnsafeFakeServiceServer interface {
	mustEmbedUnimplementedFakeServiceServer()
}

func RegisterFakeServiceServer(s grpc.ServiceRegistrar, srv FakeServiceServer) {
	s.RegisterService(&FakeService_ServiceDesc, srv)
}

func _FakeService_GetFakeRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeServiceServer).GetFakeRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FakeService_GetFakeRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeServiceServer).GetFakeRequest(ctx, req.(*FakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FakeService_ServiceDesc is the grpc.ServiceDesc for FakeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FakeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fakeapi.FakeService",
	HandlerType: (*FakeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFakeRequest",
			Handler:    _FakeService_GetFakeRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
