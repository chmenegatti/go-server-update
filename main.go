package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-server-updater/src/configuration/database/mongodb"
	"github.com/go-server-updater/src/configuration/logger"
	"github.com/go-server-updater/src/controller/routes"
)

func main() {

	logger.Info("Starting the application...")

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error while connecting to the database: %v", err.Error())
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
