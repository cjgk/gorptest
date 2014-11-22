package models

import (
	"github.com/coopernurse/gorp"
	"code.google.com/p/go.crypto/bcrypt"
    "errors"
)


type User struct {
	Id       int         `json:"id"`
	Deleted  bool        `json:"-"`
	Email    string      `json:"email"`
	Name     string      `json:"name"`
	Password string      `json:"-"`
}

type LiveUserService struct{
    Db *gorp.DbMap
}

func (s LiveUserService) Create() error { return nil }
func (s LiveUserService) Retrieve(id int) (interface{}, error) {
    user := new(User)
	obj, err := s.Db.Get(User{}, id)
	if err != nil {
		return user, err
	} else if obj == nil {
		return user, errors.New("Not found")
	}
	user = obj.(*User)

    return user, nil
}

func (s LiveUserService) Update() error { return nil }

func (s LiveUserService) Delete(id int) error {
    user, err := s.Retrieve(id)
    if err != nil {
        return err
    }

	_, err = s.Db.Delete(user)
    if err != nil {
        return nil
    }

    return nil
}

func NewUserService(dbmap *gorp.DbMap) TableService {
    return LiveUserService{Db: dbmap}
}

func NewUser(name, email, password string) (User, error) {
	pwHash, err := hashPw(password)
    user := User{}
	if err != nil {
		return User{}, err
	}

	 user = User{
		Deleted:  false,
		Email:    email,
		Name:     name,
		Password: pwHash,
	}

    return user, nil
}

func hashPw(pass string) (string, error) {
	bytePass := []byte(pass)
	pwHash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	strHashPass := string(pwHash)

	return strHashPass, nil
}

func validatePw(pass string, hash string) error {
	bytePass := []byte(pass)
	byteHash := []byte(hash)

	return bcrypt.CompareHashAndPassword(byteHash, bytePass)
}
