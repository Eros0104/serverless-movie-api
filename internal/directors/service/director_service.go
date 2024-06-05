package service

import (
	"movie-api/internal/directors/model"
)

var directors = []model.Director{
	{ID: "1", Name: "Christopher Nolan", BirthDate: "30-07-1970", Nationality: "British"},
	{ID: "2", Name: "Wachowski Brothers", BirthDate: "21-06-1965", Nationality: "American"},
}

func GetAllDirectors() ([]model.Director, error) {
	return directors, nil
}
