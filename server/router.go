package server

import (
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	fakeapi "template-api-go/example/pb/fakeapi"

	"template-api-go/server/internal/handler"
)

// setupRoutes - the root route function.
func (s *Server) setupRoutes() {
	fakeapi.RegisterFakeServiceServer(s.GrpcServer, &handler.ExampleServer{DB: s.DB})
	healthpb.RegisterHealthServer(s.GrpcServer, &handler.Health{})

}
