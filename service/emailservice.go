package service

import (
	"github.com/aws/aws-lambda-go/events"
	"studentpals/email-service/util"
)

func StudentAuthRequest(event events.APIGatewayProxyRequest) (string, util.HTTPStatus) {

	return "", util.HTTPStatus{}
}
