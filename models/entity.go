package models

import "sync"

type Entity struct {
	Id          string  `json:"imdbId"`
	Name        string  `json:"name"`
	PosterId    string  `json:"posterId"`
	PosterLink  string  `json:"posterLink"`
	Duration    string  `json:"duration"`
	ReleaseDate string  `json:"releaseDate"`
	Rating      float32 `json:"rating"`
	Type        string  `json:"type"`
	Mux         sync.Mutex
}
