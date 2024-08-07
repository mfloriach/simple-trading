package authentication

import (
	"context"

	pb "example.com/grpc-todo/gen"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authenticationServer struct {
	pb.UnimplementedAuthenticationServiceServer
	service   AuthenticationService
	validator *protovalidate.Validator
}

func NewAutheticationGrpcServer() *authenticationServer {
	validator, _ := protovalidate.New()
	return &authenticationServer{service: NewAuthenticationService(), validator: validator}
}

func (s *authenticationServer) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInReply, error) {
	if err := s.validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.service.SignIn(ctx, in.GetUsername(), in.GetPassword())

	return &pb.SignInReply{Token: token}, err
}
