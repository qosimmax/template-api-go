package server

import (
	fakeapi "template-api-go/example/pb/fakeapi"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"template-api-go/server/internal/handler"
)

// setupRoutes - the root route function.
func (s *Server) setupRoutes() {
	fakeapi.RegisterFakeServiceServer(s.GrpcServer, &handler.ExampleServer{DB: s.DB})
	healthpb.RegisterHealthServer(s.GrpcServer, &handler.Health{})

}
