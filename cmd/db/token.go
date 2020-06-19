package db

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Bearer string `json:"bearer"`
}

func (u *User) GenerateToken() JWT {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"iss":   "https://tasker.jonasburster.de",
		"exp":   time.Now().Add(8760 * time.Hour).Format(time.RFC3339),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ifPanic(err)
	}

	return JWT{Bearer: tokenString}
}

func AuthToken(tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	ifPanic(err)
	if token.Valid {
		return token
	}
	return nil
}
