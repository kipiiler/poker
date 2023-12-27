package handler

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {

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
	}
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hooray!",
	// 	})
	// })
	// err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// if err != nil {
	// 	panic(err)
	// }

	return &Router{
		Engine: router,
	}
}
