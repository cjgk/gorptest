package main

import (
	"github.com/cjgk/gorptest/controllers"
	"github.com/cjgk/gorptest/models"
	"github.com/gorilla/mux"
	"net/http"
)


func main() {
	// Set up Gorp
	dbmap := models.InitDb()
	defer dbmap.Db.Close()

    // Set up Table services
    tableServices := models.InitTableServices(dbmap)

	// Set up router
	router := mux.NewRouter()
	router.StrictSlash(false)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.HomeHandlerGet(w, r, dbmap)
	})

    users := &controllers.UserController{Service: tableServices}
	router.HandleFunc("/users", users.Action(users.Index)).Methods("GET")
	router.HandleFunc("/users/{key}", users.Action(users.Get)).Methods("GET")
	router.HandleFunc("/users", users.Action(users.Post)).Methods("POST")
	router.HandleFunc("/users/{key}", users.Action(users.Post)).Methods("PUT")
	router.HandleFunc("/users/{key}", users.Action(users.Delete)).Methods("DELETE")

	http.Handle("/", router)

	http.ListenAndServe(":3001", nil)
}
