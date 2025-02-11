package main

import (

	// "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/registry"
	// "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/attractions"
	// "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/tracing"
	// "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/tune"

	// "github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"

	"log"
	"net/http"
	
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

	// "github.com/pTpTk/HotelReservation/cmd/services/attractions"
	// "github.com/bradfitz/gomemcache/memcache"
)

func main() {
	log.Println("Starting service...")
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: "service under construction",
	}, nil
}
