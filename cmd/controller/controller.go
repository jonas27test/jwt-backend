package controller

import (
	"log"

	"github.com/jonas27test/jwt-backend/cmd/db"
)

type Controller struct {
	DB     db.DB
	Secret string
}

func ifPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
