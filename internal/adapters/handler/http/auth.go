package adapters

import (
	_ "holdem/docs"
	service "huskyholdem/service"
	user "huskyholdem/user"
	"huskyholdem/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		UserService: userService,
	}
}

func (h *AuthHandler) AuthUserWithEmail(c *gin.Context) {
	var userInfo user.User
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		utils.HandleError(c, utils.ErrBadRequest)
		return
	}

	err := h.UserService.Login(userInfo.Email, userInfo.Password)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	authToken, err := h.UserService.GenerateAuthToken(userInfo.Email)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	data := utils.NewAuthResponseMessage(authToken)
	utils.HandleSuccessWithMessage(c, data, "success")

}

func (h *AuthHandler) GenerateBotToken(c *gin.Context) {
}
