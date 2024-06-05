package handler

import (
	"context"
	"encoding/json"
	"movie-api/internal/directors/service"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func GetDirector(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	directors, err := service.GetAllDirectors()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	body, err := json.Marshal(directors)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}
