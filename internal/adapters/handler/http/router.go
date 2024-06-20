package adapters

import (
	"fmt"
	_ "holdem/docs"
	adapters "huskyholdem/adapters/handler/ws"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(
	pingHandler *PingHandler,
	authHandler *AuthHandler,
	botHandler *BotHandler,
	gameHandler *adapters.GameHandler,
) (*Router, error) {

	// Cors
	config := cors.DefaultConfig()
	allowedOrigins := os.Getenv("HTTP_ALLOWED_ORIGINS")
	originsList := strings.Split(allowedOrigins, ",")
	config.AllowOrigins = originsList
	fmt.Println("Allowed origins: ", originsList)
	fmt.Println("AAAAA")

	router := gin.Default()
	router.Use(cors.New(config))

	// Routes
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Version 1.0.0
	v1 := router.Group("/v1")
	{
		ping := v1.Group("/ping")
		ping.GET("/", pingHandler.Ping)
		ping.GET("/env", pingHandler.GetEnv)

		auth := v1.Group("/auth")
		auth.POST("/login", authHandler.AuthUserWithEmail)
		auth.POST("/bot", authHandler.GenerateBotToken)

		bot := v1.Group("/bot").Use(AuthTokenMiddleware(authHandler.UserService))
		bot.POST("/new", botHandler.CreateNewBot)
		bot.GET("/:botId", botHandler.GetBotByID)
		bot.PUT("/:botId", botHandler.UpdateBotMetadata)
		bot.POST("/:botId/generate", botHandler.GenerateBotAuthToken)

		botAuth := v1.Group("/bot/token").Use(AuthBotTokenMiddleware(botHandler.BotService))
		botAuth.GET("/self", botHandler.GetBotByToken)
		botAuth.PUT("/update", botHandler.UpdateBotMetadataByToken)
		botAuth.POST("/key", botHandler.AddKeyValueToCache)
		botAuth.GET("/key/:key", botHandler.GetBotKeyFromCache)
		botAuth.DELETE("/key/:key", botHandler.RemoveBotKeyFromCache)

		game := v1.Group("/game")
		game.POST("/new", gameHandler.CreateNewGame)
		game.GET("/ws", gameHandler.BeNice)
		game.GET("/room/:room", gameHandler.GameSocketByID)
	}

	return &Router{
		router: router,
	}, nil
}

func (r *Router) Run(address string) error {
	return r.router.Run(address)
}
