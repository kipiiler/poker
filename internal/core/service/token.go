package service

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewUserClaims(email string) UserClaims {
	return UserClaims{
		Email: email,
	}
}

func NewAuthToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("AUTH_TOKEN_SECRET")
	if secret == "" {
		return "", errors.New("AUTH_TOKEN_SECRET not set")
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseAuthAccessToken(accessToken string) (*UserClaims, error) {
	secret := os.Getenv("AUTH_TOKEN_SECRET")
	if secret == "" {
		return nil, errors.New("AUTH_TOKEN_SECRET not set")
	}
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedAccessToken.Claims.(*UserClaims), nil
}
