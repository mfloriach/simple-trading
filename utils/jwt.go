package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtAuthorization struct {
	secret []byte
}

func NewAuthenticationService() JwtAuthorization {
	return JwtAuthorization{
		secret: []byte("ertretretertret"),
	}
}

func (s JwtAuthorization) SignIn(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  username,
		"userID":    1,
		"sessionID": uuid.New().String(),
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(s.secret)
}

func (s JwtAuthorization) Verify(tokenString string) (bool, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return s.secret, nil
	})
	if err != nil {
		return false, map[string]interface{}{}, err
	}

	fmt.Println(token.Claims)

	c, ok := token.Claims.(jwt.MapClaims)
	if ok {
		return true, c, nil
	}

	return false, map[string]interface{}{}, nil
}
