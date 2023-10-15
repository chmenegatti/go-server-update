package service

import (
	"fmt"

	"github.com/go-server-updater/src/configuration/logger"
	"github.com/go-server-updater/src/configuration/rest_err"
	"github.com/go-server-updater/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("CreateUser", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
