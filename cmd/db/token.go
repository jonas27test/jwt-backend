package db

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Token string `json:"token"`
}

func (u *User) GenerateToken() JWT {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"iss":   "tasker",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ifPanic(err)
	}

	return JWT{Token: tokenString}
}
