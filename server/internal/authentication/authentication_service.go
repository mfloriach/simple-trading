package authentication

import (
	"context"

	"example.com/grpc-todo/utils"
)

type AuthenticationService struct {
	jwt utils.JwtAuthorization
}

func NewAuthenticationService() AuthenticationService {
	return AuthenticationService{
		jwt: utils.NewAuthenticationService(),
	}
}

func (s AuthenticationService) SignIn(ctx context.Context, username string, password string) (string, error) {
	return s.jwt.SignIn(username)
}

func (s AuthenticationService) Verify(token string) (bool, map[string]interface{}, error) {
	return s.jwt.Verify(token)
}
