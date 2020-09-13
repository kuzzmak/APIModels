package models

type Episode struct {
	Id          string  `json:"id"`
	SeasonId    string  `json:"seasonId"`
	Name        string  `json:"name"`
	Rating      float32 `json:"rating"`
	Description string  `json:"description"`
}
