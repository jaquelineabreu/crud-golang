package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/crud-golang/src/configuration/logger"
	"github.com/jaquelineabreu/crud-golang/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user app")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
