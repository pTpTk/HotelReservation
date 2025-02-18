package main

import (

	"log"
	"context"
	
	"github.com/aws/aws-lambda-go/lambda"

)

func main() {
	log.Println("Starting service...")
	lambda.Start(handler)
}

type UserRequest struct {
	Username string     
	Password string
}

type UserResponse struct {
	Correct bool
}

func handler(ctx context.Context, req UserRequest) (UserResponse, error) {
	return CheckUser(ctx, &req)
}
