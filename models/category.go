package models

type CategoryUpdate struct {
	UserId   int64  `json:"userId"`
	Category string `json:"category"`
}
