package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"template-api-go/client/database"
	"template-api-go/client/pubsub"
	"template-api-go/config"
	"template-api-go/monitoring/metrics"
	"template-api-go/monitoring/trace"
	fakeapi "template-api-go/proto/fake-api"
	"template-api-go/server/internal/handler"
)

// Server holds the HTTP server, router, config and all clients.
type Server struct {
	Config     *config.Config
	DB         *database.Client
	PubSub     *pubsub.Client
	HTTP       *http.Server
	Router     *mux.Router
	GrpcServer *grpc.Server
}

// Create sets up the HTTP server, router and all clients.
// Returns an error if an error occurs.
func (s *Server) Create(ctx context.Context, config *config.Config) error {
	metrics.RegisterPrometheusCollectors()

	var dbClient database.Client
	if err := dbClient.Init(ctx, config); err != nil {
		return fmt.Errorf("database client: %w", err)
	}

	var psClient pubsub.Client
	//	if err := psClient.Init(config); err != nil {
	//		return fmt.Errorf("pubsub client: %w", err)
	//	}

	s.GrpcServer = grpc.NewServer()
	s.DB = &dbClient
	s.PubSub = &psClient
	s.Config = config
	s.Router = mux.NewRouter()
	s.HTTP = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: s.Router,
	}

	s.setupRoutes()

	return nil
}

// Serve tells the server to start listening and serve HTTP requests.
// It also makes sure that the server gracefully shuts down on exit.
// Returns an error if an error occurs.
func (s *Server) Serve(ctx context.Context) error {
	closer, err := trace.InitGlobalTracer(s.Config)

	if err != nil {
		return fmt.Errorf("init global tracer: %w", err)
	}

	defer closer.Close()

	idleConnsClosed := make(chan struct{}) // this is used to signal that we can not exit
	go func(ctx context.Context, httpSrv *http.Server, grpcSrv *grpc.Server) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

		<-stop

		log.Info("Shutdown signal received")

		if err := httpSrv.Shutdown(ctx); err != nil {
			log.Error(err.Error())
		}

		grpcSrv.GracefulStop()

		close(idleConnsClosed) // call close to say we can now exit the function
	}(ctx, s.HTTP, s.GrpcServer)

	log.Infof("Ready at: %s", s.Config.Port)

	go func(grpcSrv *grpc.Server) {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8085"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		fakeapi.RegisterExampleServer(grpcSrv, &handler.ExampleServer{DB: s.DB})
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}(s.GrpcServer)

	if err := s.HTTP.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("unexpected server error: %w", err)
	}
	<-idleConnsClosed // this will block until close is called

	return nil
}
