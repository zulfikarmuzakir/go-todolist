package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func ConnectDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvConfig("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := NewMongoContext()
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(EnvConfig("MONGO_DATABASE"))

	return database
}

func GetCollection(database *mongo.Database, collectionName string) *mongo.Collection {
	collection := ConnectDB().Collection(collectionName)
	return collection
}
