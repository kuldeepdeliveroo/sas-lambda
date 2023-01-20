package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func HandleRequest() (events.APIGatewayProxyResponse, error) {
	// This lambda will be invoked by a URL with the format similar to https://roo-domain.io/2343cf3d2
	log.Printf("Hello World!!")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hey hey santa",
	}, nil
}
