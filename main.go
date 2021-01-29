package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var headers = map[string]string {
	"Content-Type": "application/json",
	"X-Custom-Header": "application/json",
}

func HandleRequest(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	go LogEnvironment(event)

	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           headers,
		MultiValueHeaders: nil,
		Body:              "",
		IsBase64Encoded:   false,
	}, nil
}

func LogEnvironment(event events.APIGatewayProxyRequest) {

}


func main() {
	lambda.Start(HandleRequest)
}
