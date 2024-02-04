package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthReponseMessage struct {
	AuthToken string `json:"authToken" example:""`
}

func NewAuthResponseMessage(authToken string) *AuthReponseMessage {
	return &AuthReponseMessage{AuthToken: authToken}
}

type AuthRequestMessage struct {
	Email    string `json:"email" example:""`
	Password string `json:"password" example:""`
}

func NewAuthRequestMessage(email string, password string) *AuthRequestMessage {
	return &AuthRequestMessage{Email: email, Password: password}
}

type EnvResponseMessage struct {
	Enviroment     string `json:"enviroment" example:"development"`
	AllowedOrigins string `json:"allowedOrigins" example:"http://localhost:3000"`
	Port           string `json:"port" example:"8080"`
	Version        string `json:"version" example:"1.0.0"`
}

func NewEnvResponseMessage(env string, allowedOrigins string, port string, version string) *EnvResponseMessage {
	return &EnvResponseMessage{Enviroment: env, AllowedOrigins: allowedOrigins, Port: port, Version: version}
}

type responseMessage struct {
	Data    any    `json:"data"`
	Message string `json:"message" example:"Success"`
	Success bool   `json:"success" example:"true"`
}

type errorResponseMessage struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// Handle success responses
func HandleSuccess(c *gin.Context, data any) {
	res := responseMessage{Data: data, Message: "success", Success: true}
	c.JSON(http.StatusOK, res)
}

// Handle success responses with message
func HandleSuccessWithMessage(c *gin.Context, data any, message string) {
	res := responseMessage{Data: data, Message: message, Success: true}
	c.JSON(http.StatusOK, res)
}

// Handle success responses without data
func HandleSuccessWithoutData(c *gin.Context, message string) {
	res := responseMessage{Data: nil, Message: message, Success: true}
	fmt.Println(res)
	c.JSON(http.StatusOK, res)
}

// Handle error responses
func HandleError(c *gin.Context, err error) {
	statusCode, ok := ErrorMapToHttp[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}
	rsp := errorResponseMessage{Message: err.Error(), Success: false}
	c.JSON(statusCode, rsp)
}
