package controllers

import (
	"log"
	"net/http"
)

func err500(err error, w http.ResponseWriter) error {
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	return err
}

func err400(err error, w http.ResponseWriter) error {
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	return err
}

func err404(w http.ResponseWriter) {
	http.Error(w, "Resource not found", http.StatusBadRequest)
}
