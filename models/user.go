package models

import (
	"time"
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
