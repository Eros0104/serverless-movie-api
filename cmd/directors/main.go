package main

import (
	"movie-api/internal/directors/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.GetDirector)
}
