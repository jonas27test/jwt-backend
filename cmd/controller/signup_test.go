package controller

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jonas27test/jwt-backend/cmd/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbURL = "mongodb://0.0.0.0:27017"
)

func Test_Signup(t *testing.T) {
	os.Setenv("SECRET", "test")
	log.SetFlags(log.Lshortfile)

	email := "jonas@test.t"
	pass := "pass"
	req, err := http.NewRequest("Post", "/signin", strings.NewReader("{\"email\":\""+email+"\", \"password\": \""+pass+"\"}"))
	tFatal(t, err)
	c := Controller{DB: db.DB{DB: db.Connection(dbURL)}}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.Signup)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != "" {
		t.Errorf("handler returned unexpected body: got %v does not contain %v",
			rr.Body.String(), email)
	}

	cleanup(c.DB.DB)
}

func cleanup(coll *mongo.Collection) {
	coll.Drop(context.Background())
}

func tFatal(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
