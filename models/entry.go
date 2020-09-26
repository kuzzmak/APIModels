package models

import "time"

type Entry struct {
	Id        string    `json:"id"`
	UserId    int64     `json:"userId"`
	Category  string    `json:"category"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	Comment   string    `json:"comment"`
}
