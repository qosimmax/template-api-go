package server

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"template-api-go/server/internal/handler"
)

const v1API string = "/template-api-go/api/v1"

// setupRoutes - the root route function.
func (s *Server) setupRoutes() {
	s.Router.Handle("/metrics", promhttp.Handler()).Name("Metrics")
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods(http.MethodGet).Name("Health")
}
