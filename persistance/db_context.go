package persistance

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lungo/config"
	"time"
)

type DbContext struct {
	Client *mongo.Client
}

func NewContext(appConfig *config.Config) *DbContext {
	client, err := mongo.NewClient(options.Client().ApplyURI(appConfig.MongoEndpoint))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return &DbContext{
		Client: client,
	}
}
