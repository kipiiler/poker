package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnvReposnseMessage struct {
	enviroment     string `json:"enviroment" example:"development"`
	allowedOrigins string `json:"allowedOrigins" example:"http://localhost:3000"`
	version        string `json:"version" example:"1.0.0"`
}

func HandleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func HandleError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
