package models

import (
	"fmt"
	"time"
)

// Entry presents one user expense entry in database.
type Entry struct {
	ID        string    `json:"ID"`
	UserID    int64     `json:"UserID"`
	Category  string    `json:"Category"`
	Amount    float32   `json:"Amount"`
	CreatedAt time.Time `json:"CreatedAt"`
	Comment   string    `json:"Comment"`
}

func (e *Entry) String() (str string) {
	return fmt.Sprintf("Entry\n\tId: %v\n\tUserId: %v\n\tCategory: %v\n\tAmount: %v\n\tCreatedAt: %v\n\tComment: %v\n",
		e.ID, e.UserID, e.Category, e.Amount, e.CreatedAt, e.Comment)
}
