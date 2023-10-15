package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-server-updater/src/configuration/logger"
	"github.com/go-server-updater/src/controller"
	"github.com/go-server-updater/src/controller/routes"
	"github.com/go-server-updater/src/model/service"
)

func main() {

	logger.Info("Starting the application...")

	svc := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(svc)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
