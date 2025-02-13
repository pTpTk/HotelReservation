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
	Reserve RequestType = iota
	Check
)

type ReservationRequest struct {
	RequestType  RequestType
	CustomerName string     
	HotelID    []string
	InDate       string      
	OutDate      string      
	RoomNumber   int         
}

type ReservationResponse struct {
	HotelID []string
}

func handler(ctx context.Context, req ReservationRequest) (ReservationResponse, error) {
	switch req.RequestType {
		case Reserve:
			log.Println("Recieved MakeReservation request")
			return MakeReservation(ctx, &req)
		case Check:
			log.Println("Recieved CheckAvailability request")
			return CheckAvailability(ctx, &req)
		default:
			return ReservationResponse{}, fmt.Errorf("unknown req type: %d", req.RequestType)
	}
}
