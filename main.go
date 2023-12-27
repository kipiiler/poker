package main

import (
	"fmt"
	handler "huskyholdem/handler/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	port := os.Getenv("PORT")
	pingHandler := handler.NewPingHandler()

	router, err := handler.NewRouter(pingHandler)
	if err != nil {
		fmt.Println("Unable to start application: " + err.Error())
		os.Exit(1)
	}

	err = router.Run(":" + port)
	if err != nil {
		fmt.Println("Unable to start http: " + err.Error())
		os.Exit(1)
	}
}
