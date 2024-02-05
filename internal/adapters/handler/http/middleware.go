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

		header := c.Request.Header["Authorization"]
		if len(header) == 0 {
			utils.HandleError(c, utils.ErrUnauthorized)
			return
		}

		token := strings.Split(header[0], " ")[1]

		userClaim, errParse := service.ParseAuthAccessToken(token)
		if errParse != nil {
			utils.HandleError(c, utils.ErrUnauthorized)
			return
		}

		isTokenValid, err := us.CheckAuthToken(userClaim.Email, token)
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
