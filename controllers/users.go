package controllers

import (
    "log"
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"
	"github.com/gorilla/mux"
    "github.com/coopernurse/gorp"
    "github.com/cjgk/gorptest/models"
)

func UserHandlerList (w http.ResponseWriter, r *http.Request, db *gorp.DbMap) {
    var (
        users []models.User
    )
    _, err := db.Select(&users, "select * from users order by id")
    if err500(err, w) != nil {
        return
    }

    log.Println(users)
    jsonUsers, err := json.Marshal(users)
    if err500(err, w) != nil {
        return
    }

    fmt.Fprint(w, string(jsonUsers))
}

func UserHandlerGet (w http.ResponseWriter, r *http.Request, db *gorp.DbMap) {
    vars := mux.Vars(r)

    userId, err := strconv.Atoi(vars["key"])
    if err400(err, w) != nil {
        return
    }

    obj, err := db.Get(models.User{}, userId)
    if err500(err, w) != nil {
        return
    } else if obj == nil {
        err404(w)
        return
    }

    user := obj.(*models.User)

    jsonUser, err := json.Marshal(user)
    if err500(err, w) != nil {
		return
	}

    fmt.Fprint(w, string(jsonUser))
}

func UserHandlerPost (w http.ResponseWriter, r *http.Request, db *gorp.DbMap) {
    name := r.FormValue("name")
    email := r.FormValue("email_address")
    password := r.FormValue("password")

    user, err := models.NewUser(name, email, password)
    if err500(err, w) != nil {
        return
    }

    err = db.Insert(&user)
    if err500(err, w) != nil {
        return
    }

    jsonUser, err := json.Marshal(user)
    if err500(err, w) != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, string(jsonUser))
}

func UserHandlerDelete (w http.ResponseWriter, r *http.Request, db *gorp.DbMap) {
    vars := mux.Vars(r)

    userId, err := strconv.Atoi(vars["key"])
    if err400(err, w) != nil {
        return
    }

    obj, err := db.Get(models.User{}, userId)
    if err500(err, w) != nil {
        return
    } else if obj == nil {
        err404(w)
        return
    }

    user := obj.(*models.User)

    _, err = db.Delete(user)
    if err500(err, w) != nil {
        return
    }
}

func UserHandlerPut (w http.ResponseWriter, r *http.Request, db *gorp.DbMap) {
    vars := mux.Vars(r)

    name := r.FormValue("name")
    email := r.FormValue("email_address")
    password := r.FormValue("password")

    userId, err := strconv.Atoi(vars["key"])
    if err400(err, w) != nil {
        return 
    }

    obj, err := db.Get(models.User{}, userId)
    if err500(err, w) != nil {
        return
    } else if obj == nil {
        err404(w)
        return
    }

    user := obj.(*models.User)

    log.Println(user)
    if len(name) > 0 {
        user.Name = name
    }

    if len(email) > 0 {
        user.Email = email
    }

    if len(password) > 0 {
        pwHash, err := models.HashPw(password)
        if err500(err, w) != nil {
            return
        }

        user.Password = pwHash
    }

    _, err = db.Update(user)
    if err500(err, w) != nil {
        return
    }

    jsonUser, err := json.Marshal(user)
    if err500(err, w) != nil {
		return
	}

    fmt.Fprint(w, string(jsonUser))
}
