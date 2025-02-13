package main

import (
	"context"

	"log"
)

// NearbyRest returns all restaurants close to the hotel.
func Nearby(ctx context.Context, req *GeoRequest) ([]string, error) {
	log.Println("In Attractions NearbyRest")
	
	res := []string {"hot_1", "hot_2", "hot_3"}

	return res, nil
}

