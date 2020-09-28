package models

import (
	"fmt"
	"time"
)

type Entry struct {
	Id        string    `json:"id"`
	UserId    int64     `json:"userId"`
	Category  string    `json:"category"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	Comment   string    `json:"comment"`
}

func (e *Entry) String() (str string) {
	return fmt.Sprintf("Entry\n\tId: %v\n\tUserId: %v\n\tCategory: %v\n\tAmount: %v\n\tCreatedAt: %v\n\tComment: %v\n",
		e.Id, e.UserId, e.Category, e.Amount, e.CreatedAt, e.Comment)
}
