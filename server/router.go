package server

import (
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	fakeapi "template-api-go/proto/fake-api"

	"template-api-go/server/internal/handler"
)

// setupRoutes - the root route function.
func (s *Server) setupRoutes() {
	fakeapi.RegisterExampleServer(s.GrpcServer, &handler.ExampleServer{DB: s.DB})
	healthpb.RegisterHealthServer(s.GrpcServer, &handler.Health{})

}
