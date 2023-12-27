package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hooray!",
	})
}

type EnvReposnseMessage struct {
	enviroment     string `json:"enviroment" example:"development"`
	allowedOrigins string `json:"allowedOrigins" example:"http://localhost:3000"`
	version        string `json:"version" example:"1.0.0"`
}

func (h *PingHandler) GetEnv(c *gin.Context) {
	allowOrigin := os.Getenv("HTTP_ALLOWED_ORIGINS")
	env := os.Getenv("ENV")
	version := os.Getenv("VERSION")
	c.JSON(http.StatusOK, gin.H{
		"message": "hooray!",
	})
}
