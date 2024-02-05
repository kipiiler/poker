package adapters

import (
	"huskyholdem/service"
	"huskyholdem/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthTokenMiddleware(us *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		type messagePayload struct {
			Email string `json:"email"`
		}

		var message messagePayload
		token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

		if err := c.ShouldBindJSON(&message); err != nil {
			utils.HandleError(c, utils.ErrBadRequest)
			return
		}

		isTokenValid, err := us.CheckAuthToken(message.Email, token)
		if err != nil {
			utils.HandleError(c, utils.ErrUnauthorized)
			return
		}
		if !isTokenValid {
			utils.HandleError(c, utils.ErrUnauthorized)
			return
		}
		c.Next()
	}
}
