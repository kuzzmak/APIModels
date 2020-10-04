package models

type MonthlyGraphDataRequest struct {
	UserId int64
	Month  int
}

type Graph struct {
	Label        string
	Labels       []string
	Data         []float32
	CurrentYear  int
	CurrentMonth int
	CurrentDay   int
	Colors       []string
	Min          float32
	Max          float32
	Months       []int
}
