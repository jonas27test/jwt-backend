package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/jonas27test/jwt-backend/cmd/db"
)

func (c *Controller) Verify(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	headerToken := strings.Split(authHeader, " ")
	if len(headerToken) == 2 {
		token := db.AuthToken(headerToken[1])
		if token != nil {
			log.Panicln("token not correct")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(token.Claims)
	}
	log.Panicln("no token")
}
