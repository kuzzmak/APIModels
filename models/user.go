package models

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
	"unicode"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
	Birthday  time.Time
}

// Validate checks if user gave valid username, first name and last name
func (u *User) Validate() (ok bool, err error) {

	// name check
	ok = checkName(u.FirstName)
	if !ok {
		return false, errors.New("first name needs to consist only of letters")
	}

	ok = checkName(u.LastName)
	if !ok {
		return false, errors.New("last name needs to consist only of letters")
	}

	exists, err := u.Exists()
	if exists {
		return false, errors.New("user exists")
	}

	return ok, nil
}

// Exists checks if user with provided userName already exists in database
func (u *User) Exists() (exists bool, err error) {

	resp, err := http.Get("http://localhost:7070/dbAPI/user/" + u.UserName)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)

	if string(html) == "false" {
		return false, nil
	}

	return true, nil
}

// checkName checks if name or last name consists only of letters
func checkName(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
