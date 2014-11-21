package controllers

import (
	"github.com/coopernurse/gorp"
	"log"
	"net/http"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func HomeHandlerGet(w http.ResponseWriter, r *http.Request, db *gorp.DbMap) {
	log.Println("HOME")
}
