package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jonas27test/jwt-backend/cmd/db"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Signin(w http.ResponseWriter, r *http.Request) {
	u := db.UserFromRequest(r)
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	ifPanic(err)
	u.Password = string(hash)

	if !c.DB.InsertUser(u) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u.GenerateToken())
}
