package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jonas27test/jwt-backend/cmd/db"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	u := db.UserFromRequest(w, r)

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	ifPanic(err)
	u.Password = string(hash)

	if c.DB.FetchUser(u.Email) != (db.User{}) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !c.DB.InsertUser(u) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
