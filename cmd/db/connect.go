package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	DB *mongo.Collection
}

func Connection(url string) *mongo.Collection {
	var client *mongo.Client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(url))
	ifPanic(err)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	ifPanic(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	ifPanic(err)
	log.Println("Connected to mongoDB!")
	return client.Database("jwt-backend").Collection("users")
}

func ifPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
