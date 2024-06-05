package handler

import (
	"context"
	"encoding/json"
	"movie-api/internal/movies/service"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func GetMovie(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	movieID := request.PathParameters["id"]
	movie, err := service.GetMovieByID(movieID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	response, err := json.Marshal(movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(response),
	}, nil
}
