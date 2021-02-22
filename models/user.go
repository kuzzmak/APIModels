package models

import (
	"time"
)

// User represents user entry in a database.
type User struct {
	ID        int64
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
	Birthday  time.Time
}
