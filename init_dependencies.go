package main

import (
	"github.com/go-server-updater/src/controller"
	"github.com/go-server-updater/src/model/repository"
	"github.com/go-server-updater/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	svc := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(svc)
}
