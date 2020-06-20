package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jonas27test/jwt-backend/cmd/db"
)

func Test_Verify(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	os.Setenv("SECRET", "test")

	// genToken
	u := db.User{Email: "test", Password: "test"}
	token := u.GenerateToken()

	req, err := http.NewRequest(http.MethodGet, "/verify", strings.NewReader(""))
	req.Header.Set("Authorization", "Bearer: "+token.Bearer)
	tFatal(t, err)
	c := Controller{DB: db.DB{DB: db.Connection(dbURL)}}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.Verify)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !(rr.Body.String() == "") {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), "id")
	}
}
