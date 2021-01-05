package driver

import (
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

var ctx = context.Background()

func connect() (*mongo.Database, error) {
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

    return client.Database("funBloggers"), nil
}

func Insert(collectionName string, insertData interface{}) error {
	db, err := connect()
    if err != nil {
        return err
    }

    _, err = db.Collection(collectionName).InsertOne(ctx, insertData)
    if err != nil {
        return err
    }

    return nil
}