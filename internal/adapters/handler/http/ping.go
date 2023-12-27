package handler

import (
	"huskyholdem/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Ping godoc
//
//	@Summary		Check server health
//	@Description	If server is healthy, it will return "hooray!"
//	@Tags			Misc
//	@Accept			json
//	@Produce		json
//	@Success		200					{string}	string	"hooray!"
//	@Router			/ping [get]
func (h *PingHandler) Ping(c *gin.Context) {
	utils.HandleSuccessWithoutData(c, "hooray!")
}

// GetEnv godoc
//
// @Summary		Get environment variables
// @Description	Return environment variables if in development mode
// @Tags		Misc
// @Accept		json
// @Produce		json
// @Success		200					{object}	utils.EnvResponseMessage		"Environment variables"
// @Failure		403					{object}	utils.errorResponseMessage	"Forbidden to access in production"
// @Router		/ping/env [get]
func (h *PingHandler) GetEnv(c *gin.Context) {
	env := os.Getenv("ENV")
	if env == "development" {
		allowOrigin := os.Getenv("HTTP_ALLOWED_ORIGINS")
		port := os.Getenv("PORT")
		data := utils.NewEnvResponseMessage(env, allowOrigin, port, "1.0.0")
		utils.HandleSuccessWithMessage(c, data, "success")
	} else {
		utils.HandleError(c, utils.ErrForbidden)
	}
}
