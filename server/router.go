package server

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	fakeapi "template-api-go/proto/fake-api"
	"template-api-go/server/internal/handler"
)

const v1API string = "/template-api-go/api/v1"

// setupRoutes - the root route function.
func (s *Server) setupRoutes() {
	s.Router.Handle("/metrics", promhttp.Handler()).Name("Metrics")
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods(http.MethodGet).Name("Health")

	// register grpc
	fakeapi.RegisterExampleServer(s.GrpcServer, &handler.ExampleServer{DB: s.DB})
}
