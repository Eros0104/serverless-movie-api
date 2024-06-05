package main

import (
	"movie-api/internal/movies/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.GetMovie)
}
