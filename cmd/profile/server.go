package main

import (
	"context"

	"log"
)

// GetProfiles returns hotel profiles for requested IDs
func GetProfiles(ctx context.Context, req *ProfileRequest) (ProfileResponse, error) {
	log.Println("In GetProfiles")

	hot := Hotel{
		ID:          "hot_1",
		Name:        "hot_name",
		PhoneNumber: "12345",
		Description: "best hotel in town",
		Address: Address{
			StreetNumber: "street_num",
			StreetName:   "street_name",
			City:         "city",
			State:        "state",
			Country:      "country",
			PostalCode:   "000-000",
			Lat:          0.75,
			Lon:          0.75,
		},
		Images:       nil,
	}
	
	res := ProfileResponse{
		Hotels: []Hotel{hot},
	}

	return res, nil
}

