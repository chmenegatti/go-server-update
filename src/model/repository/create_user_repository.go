package repository

import (
	"context"
	"os"

	"github.com/go-server-updater/src/configuration/logger"
	"github.com/go-server-updater/src/configuration/rest_err"
	"github.com/go-server-updater/src/model"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository...")

	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error when trying to get JSON value", err)
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error when trying to insert user", err)
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil

}
