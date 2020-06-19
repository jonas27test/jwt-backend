package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jonas27test/jwt-backend/cmd/db"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Signin(w http.ResponseWriter, r *http.Request) {
	u := db.UserFromRequest(w, r)

	dbUser := c.DB.FetchUser(u.Email)
	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password)) != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u.GenerateToken())
}
