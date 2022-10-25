package main

import (
	"encoding/"
	"context"

	// SDKs

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func EchoHandler(ctx context.Context, events.APIGatewayV2HTTPRequest) {

	json.Marshal(struct{
		Greeting string `json:"greeting"`
		IP string `json:"greeting"`
		UserAgent string `json:"greeting"`
	}{
		Greeting : "Hi Lambda",
		IP : e.RequestContext.HTTP.SourceIP,
		UserAgent e.Headers["user-agent"]
	})

	if err != nil {
		return events.APIGatewayV2HTTPResponse {
			StatusCode : 400,
			Body : "error marshalling"
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode : 200,
		Body : string(body),
		Headers : map[string]string{
			"Content-Type" : "application/json",
		}, nil

	}
}