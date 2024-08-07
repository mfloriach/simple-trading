package middleware

import (
	"context"
	"log/slog"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	var err error

	defer func() {
		if r := recover(); r != nil {
			slog.Error("[PANIC] %s\n\n%s", r, string(debug.Stack()))
			err = status.Errorf(codes.Internal, "%s", req)
		}
	}()

	h, err := handler(ctx, req)

	return h, err
}
