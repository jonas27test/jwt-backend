package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jonas27test/jwt-backend/cmd/controllers"
)

var Prod = flag.Bool("prod", false, "Sets many variables to testing env.")

func main() {
	dbURL := flag.String("dbURL", "mongodb://0.0.0.0:27017", "sets the urls where to connect to the db.")
	log.SetFlags(log.Lshortfile)

	c := Controller{DB: dbConnection(*dbURL)}

	http.HandleFunc("/signup", c.Signup)
	controllers.Signin()
	// http.HandleFunc("/healthz", healthz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dbConnection(url string) *mongo.Collection {
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
	return client.Database("jwt-backend").Collection("tasker-users")
}

func ifPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
