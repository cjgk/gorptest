package controllers

import (
	"net/http"
)

type Action func(rw http.ResponseWriter, r *http.Request) error

type AppController struct{}

func (c *AppController) Action(a Action) func(rw http.ResponseWriter, r *http.Request) {
    return func(rw http.ResponseWriter, r *http.Request) {
        if err := a(rw, r); err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
        }
    }
}
