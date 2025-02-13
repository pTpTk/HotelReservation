package main

import (

	"log"
	"context"
	"fmt"
	"encoding/json"
	
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
	requestType RequestType `json:"request_type"`
	hotelId string `json:"hotel_id"`
}

func handler(ctx context.Context, event json.RawMessage) ([]string, error) {
	var attrReq AttractionsRequest
	if err := json.Unmarshal(event, &attrReq); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return nil, err
	}

	switch attrReq.requestType {
		case Rest:
			log.Println("Recieved NearbyRest request")
			return NearbyRest(ctx, &attrReq)
		case Mus:
			log.Println("Recieved NearbyMus request")
			return NearbyMus(ctx, &attrReq)
		case Cinema:
			log.Println("Recieved NearbyCinema request")
			return NearbyCinema(ctx, &attrReq)
		default:
			return nil, fmt.Errorf("unknown req type: %d", attrReq.requestType)
	}
}
