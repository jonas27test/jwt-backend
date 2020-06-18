package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jonas27test/jwt-backend/cmd/controller"
	"github.com/jonas27test/jwt-backend/cmd/db"
)

func main() {
	dbURL := flag.String("dbURL", "mongodb://0.0.0.0:27017", "sets the urls where to connect to the db.")
	log.SetFlags(log.Lshortfile)

	c := controller.Controller{DB: db.DB{DB: db.Connection(*dbURL)}}

	http.HandleFunc("/signup", c.Signup)
	// http.HandleFunc("/healthz", healthz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
