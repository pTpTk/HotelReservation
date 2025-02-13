package main

import (

	"log"
	"context"
	"fmt"
	
	"github.com/aws/aws-lambda-go/lambda"

)

func main() {
	log.Println("Starting service...")
	lambda.Start(handler)
}

// The go equivalent of enum
type RequestType int

const (
	Rest RequestType = iota
	Mus
	Cinema
)

type AttractionsRequest struct {
	RequestType RequestType
	HotelID     string
}

type AttractionsResponse struct {
	AttractionIDs []string
}

func handler(ctx context.Context, req AttractionsRequest) (AttractionsResponse, error) {

	switch req.RequestType {
		case Rest:
			log.Println("Recieved NearbyRest request")
			return NearbyRest(ctx, &req)
		case Mus:
			log.Println("Recieved NearbyMus request")
			return NearbyMus(ctx, &req)
		case Cinema:
			log.Println("Recieved NearbyCinema request")
			return NearbyCinema(ctx, &req)
		default:
			return AttractionsResponse{}, fmt.Errorf("unknown req type: %d", req.RequestType)
	}
}
