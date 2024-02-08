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

// AuthUserWithEmail godoc
//
// @Summary		Authenticate user with email and password
// @Description	Authenticate user with email and password
// @Tags			Auth
// @Accept		json
// @Produce		json
// @Param		email	body	string	true	"User Email"
// @Param		password	body	string	true	"User Password"
// @Success		200			{object}	utils.AuthResponseMessage	"Auth token"
// @Failure		400			{object}	utils.errorResponseMessage "Bad request"
// @Failure		401			{object}	utils.errorResponseMessage "Unauthorized"
// @Router		/auth/login [post]
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
