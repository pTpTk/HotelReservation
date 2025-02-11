package main

import (

	"log"
	"context"
	"net/http"
	"encoding/json"
	
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

)

func main() {
	log.Println("Starting service...")
	lambda.Start(handler)
}

// The go equivalent of enum
type RequestType int

const (
	NearbyRest RequestType = iota
	NearbyMus
	NearbyCinema
)

type AttractionsRequest struct {
	requestType RequestType `json:"request_type"`
	hotelId string `json:"hotel_id"`
}

func handler(ctx context.Context, event json.RawMessage) (events.APIGatewayProxyResponse, error) {
	var attrReq AttractionsRequest
	if err := json.Unmarshal(event, &attrReq); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return events.APIGatewayProxyResponse{}, nil
	}

	switch attrReq.requestType {
		case NearbyRest:
			log.Println("Recieved NearbyRest request")
		case NearbyMus:
			log.Println("Recieved NearbyMus request")
		case NearbyCinema:
			log.Println("Recieved NearbyCinema request")
		default:
			log.Printf("unknown req type: %d", attrReq.requestType)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: "service under construction",
	}, nil
}
