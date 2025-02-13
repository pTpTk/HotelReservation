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

type ProfileRequest struct {
	Lat float64
	Lon float64
}

type Address struct {
	StreetNumber string
	StreetName   string
	City         string
	State        string
	Country      string
	PostalCode   string
	Lat          float64
	Lon          float64
}

type Image struct {
	Url string
	Def bool
}

type Hotel struct {
	ID          string
	Name        string
	PhoneNumber string
	Description string
	Address     Address
	Images      []Image
}

type ProfileResponse struct {
	Hotels []Hotel
}

func handler(ctx context.Context, req ProfileRequest) (ProfileResponse, error) {

	return GetProfiles(ctx, &req)
}
