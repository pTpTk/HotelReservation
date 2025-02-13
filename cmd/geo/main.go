package main

import (

	"log"
	"context"
	"encoding/json"
	
	"github.com/aws/aws-lambda-go/lambda"

)

func main() {
	log.Println("Starting service...")
	lambda.Start(handler)
}

type GeoRequest struct {
	lat float64 `json:"lat"`
	lon float64 `json:"lon"`
}

func handler(ctx context.Context, event json.RawMessage) ([]string, error) {
	var geoReq GeoRequest
	if err := json.Unmarshal(event, &geoReq); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return nil, err
	}

	return Nearby(ctx, &geoReq)
}
