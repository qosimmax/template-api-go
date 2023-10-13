package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"template-api-go/client/database"
	"template-api-go/client/pubsub"
	"template-api-go/config"
	"template-api-go/monitoring/trace"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

// Server holds the HTTP server, router, config and all clients.
type Server struct {
	Config         *config.Config
	DB             *database.Client
	PubSub         *pubsub.Client
	GrpcServer     *grpc.Server
	TracerProvider *tracesdk.TracerProvider
}

// Create sets up the HTTP server, router and all clients.
// Returns an error if an error occurs.
func (s *Server) Create(ctx context.Context, config *config.Config) error {
	var dbClient database.Client
	if err := dbClient.Init(ctx, config); err != nil {
		return fmt.Errorf("database client: %w", err)
	}

	var psClient pubsub.Client
	if err := psClient.Init(config); err != nil {
		return fmt.Errorf("pubsub client: %w", err)
	}

	s.DB = &dbClient
	s.PubSub = &psClient
	s.Config = config
	s.GrpcServer = grpc.NewServer(
		//grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)

	s.setupRoutes()

	return nil
}

// Serve tells the server to start listening and serve GRPC requests.
// It also makes sure that the server gracefully shuts down on exit.
// Returns an error if an error occurs.
func (s *Server) Serve(ctx context.Context) error {
	var err error
	s.TracerProvider, err = trace.TracerProvider(s.Config)
	if err != nil {
		return fmt.Errorf("init global tracer: %w", err)
	}

	idleConnsClosed := make(chan struct{}) // this is used to signal that we can not exit
	go func(ctx context.Context) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

		<-stop

		log.Info("Shutdown signal received")
		s.shutdown(ctx)

		close(idleConnsClosed) // call close to say we can now exit the function
	}(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Service Ready at %v", lis.Addr())
	if err := s.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	<-idleConnsClosed // this will block until close is called

	return nil
}

func (s *Server) shutdown(ctx context.Context) {
	s.GrpcServer.GracefulStop()

	if err := s.TracerProvider.Shutdown(ctx); err != nil {
		log.Error(err.Error())
	}

	if err := s.DB.Close(); err != nil {
		log.Error(err.Error())
	}

}
