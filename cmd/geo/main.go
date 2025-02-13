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

type GeoRequest struct {
	lat float64
	lon float64
}

type GeoResponse struct {
	HotelIDs []string
}

func handler(ctx context.Context, req GeoRequest) (GeoResponse, error) {

	return Nearby(ctx, &req)
}
