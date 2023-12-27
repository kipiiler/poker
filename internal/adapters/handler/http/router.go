package handler

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(
	pingHandler *PingHandler,
) (*Router, error) {

	// Cors
	config := cors.DefaultConfig()
	allowedOrigins := os.Getenv("HTTP_ALLOWED_ORIGINS")
	originsList := strings.Split(allowedOrigins, ",")
	config.AllowOrigins = originsList

	router := gin.Default()
	router.Use(cors.New(config))

	// Routes

	// Version 1.0.0
	v1 := router.Group("/v1")
	{
		ping := v1.Group("/ping")
		ping.GET("/ping", pingHandler.Ping)
		ping.GET("/env", pingHandler.GetEnv)
	}

	return &Router{
		router: router,
	}, nil
}

func (r *Router) Run(address string) error {
	return r.router.Run(address)
}
