package main

import (
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	DB   *mongo.Collection
	User User
}

func writeErr(w http.ResponseWriter, status int, err string) {

}

func (c Controller) Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("hallo")
	var err error
	c.User, err = FromRequest(r)

	hash, err := bcrypt.GenerateFromPassword([]byte(c.User.Password), 10)
	ifPanic(err)

	c.User.Password = string(hash)

	// userRepo := userRepository.UserRepository{}
	// c.User = db.Signup(user)

	// if err != nil {
	// 	error.Message = "Server error."
	// 	utils.RespondWithError(w, http.StatusInternalServerError, error)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(c.User)
}
