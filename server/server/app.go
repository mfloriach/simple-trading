package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "example.com/grpc-todo/gen"
	"example.com/grpc-todo/server/internal/authentication"
	"example.com/grpc-todo/server/internal/graphs"
	"example.com/grpc-todo/server/internal/trade"
	"example.com/grpc-todo/server/server/middleware"
	"example.com/grpc-todo/utils"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// var kaep = keepalive.EnforcementPolicy{
// 	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
// 	PermitWithoutStream: true,            // Allow pings even when there are no active streams
// }

// var kasp = keepalive.ServerParameters{
// 	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
// 	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
// 	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
// 	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
// 	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
// }

func StartServer(ctx context.Context) {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("failed to listen on port 5050: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.RecoveryInterceptor,
			// middleware.AuthorizationInterceptor,
			middleware.LoggingInterceptor,
		),
		// grpc.KeepaliveEnforcementPolicy(kaep),
		// grpc.KeepaliveParams(kasp),
	)

	if err := utils.GetTimeLineDb(ctx); err != nil {
		log.Fatal(err)
	}

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("failed to listen on port 5050: %v", err)
	}

	tradeRepo := trade.NewTradeAdminCacheRepository()
	candleRepo := graphs.NewTimescaleRepo()
	pool := graphs.NewApiPool(ctx, []string{"AAPL"}, tradeRepo, candleRepo)

	reflection.Register(s)
	pb.RegisterGraphServiceServer(s, graphs.NewGrpcServer(validator, pool))
	pb.RegisterTradeServiceServer(s, trade.NewTradeGrpcServer())
	pb.RegisterAuthenticationServiceServer(s, authentication.NewAutheticationGrpcServer())
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	go func() {
		fmt.Printf("Starting Grpc server at %v\n", lis.Addr())
		if err := s.Serve(lis); err != nil && err != http.ErrServerClosed {
			fmt.Printf("server error: %s\n", err)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Shutting down server gracefully...")
		utils.TimeLineDb.Close()
		s.GracefulStop()
	}

	fmt.Println("Server stopped.")
}
