package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	dbURL = "mongodb://0.0.0.0:27017"
)

func Test_Signup(t *testing.T) {
	email := "jonas@test.t"
	req, err := http.NewRequest("Post", "/signup", strings.NewReader("{\"email\":\""+email+"\", \"email\": \"jonas@web.de\"}"))
	tFatal(t, err)
	c := Controller{DB: dbConnection(dbURL)}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
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
}

func tFatal(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
