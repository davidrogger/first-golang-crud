package main

import (
	"log"

	"github.com/davidrogger/first-golang-crud/src/configurations/logger"
	"github.com/davidrogger/first-golang-crud/src/controllers/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Initializing application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":3002"); err != nil {
		log.Fatal(err)
	}

}
