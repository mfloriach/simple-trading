package utils

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func GetMongoClient() *mongo.Database {
	if client != nil {
		return client.Database(os.Getenv("DB_NAME"))
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_CONN")))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return client.Database(os.Getenv("DB_NAME"))
}

func CloseClient() {
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}
}
