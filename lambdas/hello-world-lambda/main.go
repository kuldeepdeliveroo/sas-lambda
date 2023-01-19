package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"sas-lambda/lambdas/hello-world-lambda/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
