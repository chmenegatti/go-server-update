package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-server-updater/src/model/service"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserById(*gin.Context)
	FindUserByEmail(*gin.Context)

	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
