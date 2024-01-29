package handler

import (
	_ "holdem/docs"
	service "huskyholdem/service"
	"huskyholdem/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserService *service.UserService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) AuthUserWithEmail(c *gin.Context) {
	utils.HandleSuccessWithoutData(c, "email competed!")
}

func (h *AuthHandler) GenerateBotToken(c *gin.Context) {
}
