package db

import (
	"context"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	dbURL = "mongodb://0.0.0.0:27017"
)

func Test_RetrieveOID(t *testing.T) {
	db := DB{DB: Connection(dbURL)}
	email := "jonas@test.t"
	pass := "pass"
	u := User{Email: email, Password: pass}
	db.InsertUser(u)
	filter := bson.D{{"email", u.Email}}
	var oid OID
	_ = db.DB.FindOne(context.Background(), filter).Decode(&oid)
	if strings.Contains(oid.OID.Hex(), "00000000") {
		t.Error("contains too many zeros, probably not able to retrieve user")
	}
	db.DB.Drop(context.Background())
}
