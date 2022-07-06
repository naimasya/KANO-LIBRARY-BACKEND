package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	ctx := context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("kanolibrary"), nil
}
