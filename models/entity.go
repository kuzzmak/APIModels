package models

import "sync"

// Entity presents movie or series base.
type Entity struct {
	ID          string  `json:"ID"`
	Name        string  `json:"Name"`
	PosterID    string  `json:"PosterID"`
	PosterLink  string  `json:"PosterLink"`
	Duration    string  `json:"Duration"`
	ReleaseDate string  `json:"ReleaseDate"`
	Rating      float32 `json:"Rating"`
	Type        string  `json:"Type"`
	Mux         sync.Mutex
}
