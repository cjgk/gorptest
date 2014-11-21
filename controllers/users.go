package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/cjgk/gorptest/models"
	"github.com/coopernurse/gorp"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
    "errors"
)

type UserController struct {
    AppController
    Db *gorp.DbMap
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) error {
	var users []models.User

	_, err := c.Db.Select(&users, "select * from users order by id")
    if err != nil {
		return err
	}

	jsonUsers, err := json.Marshal(users)
    if err != nil {
		return err
	}

	fmt.Fprint(w, string(jsonUsers))

    return nil
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["key"])
	if err != nil {
		return err
	}

	obj, err := c.Db.Get(models.User{}, userId)
	if err != nil {
		return err
	} else if obj == nil {
		return errors.New("Not found")
	}

	user := obj.(*models.User)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(jsonUser))

    return nil
}

func (c *UserController) Post(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email_address")
	password := r.FormValue("password")

	user, err := models.NewUser(name, email, password)
	if err != nil {
		return err
	}

	err = c.Db.Insert(&user)
	if err != nil {
		return err
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonUser))

    return nil
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["key"])
	if err != nil {
		return err
	}

	obj, err := c.Db.Get(models.User{}, userId)
	if err != nil {
		return err
	} else if obj == nil {
		return errors.New("Not found")
	}

	user := obj.(*models.User)

	_, err = c.Db.Delete(user)
	if err != nil {
		return err
	}

    return nil
}

func (c *UserController) Put(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	name := r.FormValue("name")
	email := r.FormValue("email_address")
	password := r.FormValue("password")

	userId, err := strconv.Atoi(vars["key"])
	if err != nil {
		return err
	}

	obj, err := c.Db.Get(models.User{}, userId)
	if err != nil {
		return err
	} else if obj == nil {
		return errors.New("Not Found")
	}

	user := obj.(*models.User)

	if len(name) > 0 {
		user.Name = name
	}

	if len(email) > 0 {
		user.Email = email
	}

	if len(password) > 0 {
		pwHash, err := models.HashPw(password)
		if err != nil {
			return err
		}

		user.Password = pwHash
	}

	_, err = c.Db.Update(user)
	if err != nil {
		return err
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(jsonUser))

    return nil
}
