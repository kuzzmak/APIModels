package models

type Episode struct {
	Id          string  `json:"id"`
	SeriesId    string  `json:"seriesId"`
	Season      string  `json:"season"`
	EpNum       string  `json:"epNum"`
	Name        string  `json:"name"`
	Rating      float32 `json:"rating"`
	Description string  `json:"description"`
	PosterId    string  `json:"posterId"`
	PosterLink  string  `json:"posterLink"`
}
