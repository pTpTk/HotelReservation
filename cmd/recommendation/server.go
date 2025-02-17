package main

import (
	"context"
	"log"
)

// GetRecommendations returns recommended hotels for a given requirement
func GetRecommendations(ctx context.Context, req *RecommendationRequest) (RecommendationResponse, error) {
	log.Println("In Attractions GetRecommendations")
	
	res := RecommendationResponse{
		HotelIDs: []string {"hot_1", "hot_2", "hot_3"},
	}

	return res, nil
}

