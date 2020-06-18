package controller

import (
	"log"
	"net/http"
)

func writeErr(w http.ResponseWriter, status int, err string) {

}

func ifPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
