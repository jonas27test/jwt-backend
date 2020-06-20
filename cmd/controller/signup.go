package controller

import (
	"log"
	"net/http"

	"github.com/jonas27test/jwt-backend/cmd/db"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("signup")
	u := db.UserFromRequest(w, r)

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	ifPanic(err)
	u.Password = string(hash)

	if c.DB.RetrieveUser(u.Email) != (db.User{}) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !c.DB.InsertUser(u) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
