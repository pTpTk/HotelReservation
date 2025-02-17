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

type RateRequest struct {
	HotelIDs    []string
	InDate       string      
	OutDate      string      
}

type RoomType struct {
	BookableRate       float64
	TotalRate          float64
	TotalRateInclusive float64
	Code               string
	Currency           string
	RoomDescription    string
}

type RatePlan struct {
	HotelID  string
	Code     string
	InDate   string
	OutDate  string
	RoomType RoomType
}

type RateResponse struct {
	RatePlans []RatePlan
}

func handler(ctx context.Context, req RateRequest) (RateResponse, error) {

	return GetRates(ctx, &req)
}
