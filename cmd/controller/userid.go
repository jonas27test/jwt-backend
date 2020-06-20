package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jonas27test/jwt-backend/cmd/db"
)

func (c *Controller) UserID(w http.ResponseWriter, r *http.Request) {
	log.Println("userID")
	if os.Getenv("PROD") == "true" && !strings.Contains(r.RequestURI, "tasker.jwt-backend") {
		log.Panicln("called from outside cluster")
	}
	authHeader := r.Header.Get("Authorization")
	headerToken := strings.Split(authHeader, " ")
	if len(headerToken) == 2 {
		claims := db.AuthToken(headerToken[1])
		if claims == nil {
			log.Panicln("token not correct")
		}

		id := db.ID{ID: c.DB.RetrieveID(claims["email"].(string))}
		id.ID = "t" + id.ID + "t"
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(id)
		return
	}
	log.Panicln("no token")
}
