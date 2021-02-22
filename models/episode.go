package models

// Episode struct presents one episode in a season.
type Episode struct {
	ID          string  `json:"ID"`
	SeriesID    string  `json:"SeriesID"`
	Season      string  `json:"Season"`
	Num         string  `json:"Num"`
	Name        string  `json:"Name"`
	Rating      float32 `json:"Rating"`
	Description string  `json:"Description"`
	PosterID    string  `json:"PosterID"`
	PosterLink  string  `json:"PosterLink"`
}
