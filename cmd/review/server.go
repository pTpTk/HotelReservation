package main

import (
	"context"
	"log"
)

// GetReviewss returns hotel profiles for requested IDs
func GetReviews(ctx context.Context, req *ReviewRequest) (ReviewResponse, error) {
	log.Println("In GetReviews")

	review := ReviewComm{
		ReviewID:    "00035",
		HotelID:     "ididid",
		Name:        "good good study",
		Rating:		 4.65,
		Description: "not good",
		Images:      nil,
	}
	
	res := ReviewResponse{
		Reviews: []ReviewComm{review},
	}

	return res, nil
}

