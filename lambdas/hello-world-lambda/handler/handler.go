package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func HandleRequest(ctx context.Context, request events.ALBTargetGroupRequest) error {
	// This lambda will be invoked by a URL with the format similar to https://roo-domain.io/2343cf3d2
	log.Printf("Hello World!!")
	return nil
}
