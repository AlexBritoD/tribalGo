package main

import (
	"log"

	"search/song/api"
	"search/song/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	router := gin.Default()

	router.Use(api.AuthMiddleware())

	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/search", api.Search)
	}

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
