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

type SearchRequest struct {
	Lat     float64     
	Lon     float64
	InDate  string      
	OutDate string      
}

type SearchResponse struct {
	HotelIDs []string
}

func handler(ctx context.Context, req SearchRequest) (SearchResponse, error) {
	return Nearby(ctx, &req)
}
