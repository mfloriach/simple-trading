package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	h, err := handler(ctx, req)

	log.Println("ACCESS LOG", "Method", info.FullMethod, "Duration", time.Since(start), "Error", err, "UserID", ctx.Value("userID"), "SessionID", ctx.Value("sessionID"))

	return h, err
}
