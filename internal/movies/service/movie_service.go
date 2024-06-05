package service

import (
	"errors"
	"movie-api/internal/movies/model"
)

var movies = []model.Movie{
	{ID: "1", Title: "Inception", Year: 2010, Genre: "Sci-Fi", Director: "Christopher Nolan"},
	{ID: "2", Title: "The Matrix", Year: 1999, Genre: "Sci-Fi", Director: "Wachowski Brothers"},
}

func GetAllMovies() ([]model.Movie, error) {
	return movies, nil
}

func GetMovieByID(id string) (*model.Movie, error) {
	for _, movie := range movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, errors.New("movie not found")
}
