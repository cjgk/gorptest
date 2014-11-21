package models

import (
	"code.google.com/p/go.crypto/bcrypt"
)

type User struct {
    Id       int64  `json:"id"`
    Deleted  bool   `json:"-"`
    Email    string `json:"email"`
    Name     string `json:"name"`
    Password string `json:"-"`
}

func NewUser(name, email, password string) (User, error) {
    pwHash, err := HashPw(password)
    if err != nil {
        return User{}, err
    }

    return User {
        Deleted: false,
        Email: email,
        Name: name,
        Password: pwHash,
    }, nil
}

func HashPw(pass string) (string, error) {
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
