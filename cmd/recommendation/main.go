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

type RecommendationRequest struct {
	Require string
	Lat     float64 
	Lon     float64 
}

type RecommendationResponse struct {
	HotelIDs []string
}

func handler(ctx context.Context, req RecommendationRequest) (RecommendationResponse, error) {

	return GetRecommendations(ctx, &req)
}
