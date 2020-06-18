package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FromRequest(r *http.Request) (User, error) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Panicln(err)
		return user, nil
	}

	if user.Email == "" {
		// w.Write([]byte("Email is missing."))
		// utils.RespondWithError(w, http.StatusBadRequest, error)
		// return
	}

	if user.Password == "" {
		// error.Message = "Password is missing."
		// utils.RespondWithError(w, http.StatusBadRequest, error)
		// return
	}
	return user, nil
}
