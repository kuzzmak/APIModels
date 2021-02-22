package models

// Series represents one series which consists
// of multple season and each season has number
// of episodes.
type Series struct {
	Entity  `json:"Entity"`
	Seasons []*Season `json:"Seasons"`
}
