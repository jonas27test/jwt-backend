package db

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OID struct {
	OID primitive.ObjectID `bson:"_id"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (db *DB) InsertUser(u User) bool {
	if db.RetrieveUser(u.Email) == (User{}) {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err := db.DB.InsertOne(ctx, u)
		ifPanic(err)
		return true
	}
	return false
}

func (db *DB) RetrieveUser(email string) User {
	var u User
	filter := bson.D{{"email", email}}
	err := db.DB.FindOne(context.Background(), filter).Decode(&u)
	if err != nil {
		return User{}
	}
	return u
}

func (db *DB) RetrieveID(email string) string {
	var oid OID
	filter := bson.D{{"email", email}}
	err := db.DB.FindOne(context.Background(), filter).Decode(&oid)
	if err != nil {
		return ""
	}
	return oid.OID.Hex()
}

func (u *User) JsonString() string {
	b, err := json.Marshal(u)
	ifPanic(err)
	return string(b)
}

func UserFromRequest(w http.ResponseWriter, r *http.Request) User {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Panicln(err)
		return u
	}
	if u.Email == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Panic("Email not set")
	} else if u.Password == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Panic("Password not set")
	}
	return u
}

//
