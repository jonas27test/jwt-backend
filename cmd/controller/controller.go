package controller

import "github.com/jonas27test/jwt-backend/cmd/db"

type Controller struct {
	DB     db.DB
	Secret string
}
