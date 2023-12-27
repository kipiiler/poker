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

func (h *PingHandler) Ping(c *gin.Context) {
	utils.HandleSuccessWithoutData(c, "hooray!")
}

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
