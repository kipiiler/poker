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

type BotClaims struct {
	Email string `json:"email"`
	BotId string `json:"botId"`
	jwt.StandardClaims
}

func NewBotClaim(email string, botId string) BotClaims {
	return BotClaims{
		Email: email,
		BotId: botId,
	}
}

func NewBotToken(claims BotClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("BOT_TOKEN_SECRET")
	if secret == "" {
		return "", errors.New("BOT_TOKEN_SECRET not set")
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
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

func ParseBotToken(botToken string) (*BotClaims, error) {
	secret := os.Getenv("BOT_TOKEN_SECRET")
	if secret == "" {
		return nil, errors.New("BOT_TOKEN_SECRET not set")
	}

	parsedBotToken, err := jwt.ParseWithClaims(botToken, &BotClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedBotToken.Claims.(*BotClaims), nil
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
