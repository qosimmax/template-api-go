// Package handler contains GRPC handlers.

package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Health is used for our readiness and liveness probes.
type Health struct{}

func (h *Health) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (h *Health) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
