package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/jonas27test/jwt-backend/cmd/db"
)

func (c *Controller) Verify(w http.ResponseWriter, r *http.Request) {
	log.Println("verify")
	authHeader := r.Header.Get("Authorization")
	headerToken := strings.Split(authHeader, " ")
	if len(headerToken) == 2 {
		claims := db.AuthToken(headerToken[1])
		if claims == nil {
			log.Panicln("token not correct")
		}

		w.WriteHeader(http.StatusOK)
		return
	}
	log.Panicln("no token")
}
