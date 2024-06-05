package model

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
}
