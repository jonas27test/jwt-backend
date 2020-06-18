package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jonas27test/jwt-backend/cmd/db"
)

func Test_Signin(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	email := "jonas@test.t"
	req, err := http.NewRequest("Post", "/signin", strings.NewReader("{\"email\":\""+email+"\", \"password\": \"pass\"}"))
	tFatal(t, err)
	c := Controller{DB: db.DB{DB: db.Connection(dbURL)}}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.Signup)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if strings.Contains(rr.Body.String(), email) {
		t.Errorf("handler returned unexpected body: got %v does not contain %v",
			rr.Body.String(), email)
	}

	cleanup(c.DB.DB)
}
