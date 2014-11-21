package main

import (
	"github.com/cjgk/gorptest/controllers"
	"github.com/cjgk/gorptest/models"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Set up Gorp
	db := models.InitDb()
	defer db.Db.Close()

	// Set up router
	router := mux.NewRouter()
	router.StrictSlash(false)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.HomeHandlerGet(w, r, db)
	})

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.UserHandlerList(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/users/{key}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UserHandlerGet(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.UserHandlerPost(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/users/{key}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UserHandlerPut(w, r, db)
	}).Methods("PUT")

	router.HandleFunc("/users/{key}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UserHandlerDelete(w, r, db)
	}).Methods("DELETE")

	http.Handle("/", router)

	http.ListenAndServe(":3001", nil)
}
