package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jonas27test/jwt-backend/cmd/controller"
	"github.com/jonas27test/jwt-backend/cmd/db"
	"github.com/jonas27test/jwt-backend/cmd/kube"
)

func main() {
	dbURL := flag.String("dburl", "mongodb://0.0.0.0:27017", "sets the url where to connect to the db.")
	port := flag.String("p", ":8000", "sets the port. Use with colon, e.g. :8080")
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	log.Println("The port is: " + *port)
	log.Println("The database is on: " + *dbURL)
	// os.Setenv("SECRET", "test")

	c := controller.Controller{DB: db.DB{DB: db.Connection(*dbURL)}}
	http.HandleFunc("/userid", c.UserID)
	http.HandleFunc("/verify", c.Verify)
	http.HandleFunc("/signin", c.Signin)
	http.HandleFunc("/signup", c.Signup)
	http.HandleFunc("/healthz", kube.Healthz)
	log.Fatal(http.ListenAndServe(*port, nil))
}
