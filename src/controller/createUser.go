package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-server-updater/src/configuration/logger"
	"github.com/go-server-updater/src/configuration/validation"
	"github.com/go-server-updater/src/controller/model/request"
	"github.com/go-server-updater/src/model"
	"github.com/go-server-updater/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Create user...", zap.String("request", "POST"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error while binding JSON", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		logger.Error("Error while creating user", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Created user...", zap.String("args", fmt.Sprintf("%v", view.ConvertDomainToResponse(domain))))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
