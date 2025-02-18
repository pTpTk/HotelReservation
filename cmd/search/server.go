package main

import (
	"context"
	"log"
)

func Nearby(ctx context.Context, req *SearchRequest) (SearchResponse, error) {
	log.Println("In Search Nearby")
	
	res := SearchResponse{
		HotelIDs: []string {"rest_1", "rest_2", "rest_3"},
	}

	return res, nil
}

