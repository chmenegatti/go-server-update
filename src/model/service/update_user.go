package service

import (
	"github.com/go-server-updater/src/configuration/rest_err"
	"github.com/go-server-updater/src/model"
)

func (*userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}
