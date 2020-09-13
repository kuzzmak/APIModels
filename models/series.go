package models

type Series struct {
	Entity  `json:"entity"`
	Seasons []*Season `json:"seasons"`
}
