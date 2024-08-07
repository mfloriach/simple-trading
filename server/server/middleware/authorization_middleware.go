package middleware

import (
	"context"

	"example.com/grpc-todo/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var JwtAuthorization = utils.NewAuthenticationService()

func AuthorizationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	if info.FullMethod == "/authentication.AuthenticationService/SignIn" {
		handler(ctx, req)
	}

	if len(md["authorization"]) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Authorization token is not supplied")
	}

	ok, claims, err := JwtAuthorization.Verify(md["authorization"][0])
	if err != nil || !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Authorization token not valid")
	}

	ctx = context.WithValue(ctx, "userID", 1)
	ctx = context.WithValue(ctx, "sessionID", claims["sessionID"])

	return handler(ctx, req)
}
