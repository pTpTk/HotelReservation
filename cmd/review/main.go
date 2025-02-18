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

type ReviewRequest struct {
	HotelID string
}

type Image struct {
	Url string
	Def bool
}

type ReviewComm struct {
	ReviewID    string
	HotelID     string
	Name        string
	Rating      float64
	Description string
	Images      []Image
}

type ReviewResponse struct {
	Reviews []ReviewComm
}

func handler(ctx context.Context, req ReviewRequest) (ReviewResponse, error) {

	return GetReviews(ctx, &req)
}
