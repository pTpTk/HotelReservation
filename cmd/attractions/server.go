package main

import (
	"context"

	"log"
)

// NearbyRest returns all restaurants close to the hotel.
func NearbyRest(ctx context.Context, req *AttractionsRequest) (AttractionsResponse, error) {
	log.Println("In Attractions NearbyRest")
	
	res := AttractionsResponse{
		[]string {"rest_1", "rest_2", "rest_3"},
	}

	return res, nil
}

// NearbyMus returns all museums close to the hotel.
func NearbyMus(ctx context.Context, req *AttractionsRequest) (AttractionsResponse, error) {
	log.Println("In Attractions NearbyMus")
	
	res := AttractionsResponse{
		[]string {"mus_1", "mus_2", "mus_3"},
	}

	return res, nil
}

// NearbyCinema returns all cinemas close to the hotel.
func NearbyCinema(ctx context.Context, req *AttractionsRequest) (AttractionsResponse, error) {
	log.Println("In Attractions NearbyCinema")

	res := AttractionsResponse{
		AttractionIDs: []string{"cin_1", "cin_2", "cin_3"},
	}

	return res, nil
}

