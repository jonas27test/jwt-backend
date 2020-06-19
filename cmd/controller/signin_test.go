package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jonas27test/jwt-backend/cmd/db"
	"golang.org/x/crypto/bcrypt"
)

func Test_Signin(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	email := "jonas@test.t"
	pass := "pass"
	req, err := http.NewRequest("Post", "/signin", strings.NewReader("{\"email\":\""+email+"\", \"password\": \""+pass+"\"}"))
	tFatal(t, err)
	c := Controller{DB: db.DB{DB: db.Connection(dbURL)}}

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	ifPanic(err)
	u := db.User{Email: email, Password: string(hash)}
	c.DB.InsertUser(u)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.Signin)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "token") {
		t.Errorf("handler returned unexpected body: got %v does not contain %v",
			rr.Body.String(), email)
	}

	cleanup(c.DB.DB)
}
