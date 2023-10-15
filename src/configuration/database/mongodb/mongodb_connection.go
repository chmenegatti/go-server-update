package mongodb

import (
	"context"
	"os"

	"github.com/go-server-updater/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL    = "MONGODB_URL"
	MONGODB_DBNAME = "MONGODB_DBNAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbUri := os.Getenv(MONGODB_URL)
	mongodbDB := os.Getenv(MONGODB_DBNAME)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB!")

	return client.Database(mongodbDB), nil
}
