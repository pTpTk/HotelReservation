package main

import (
	"context"

	"log"
)

func MakeReservation(ctx context.Context, req *ReservationRequest) (ReservationResponse, error) {
	log.Println("In Reservation MakeReservation")
	log.Printf("hotel 1: %s", req.InDate)
	
	res := ReservationResponse{
		HotelID: []string {"rest_1", "rest_2", "rest_3"},
	}

	return res, nil
}

func CheckAvailability(ctx context.Context, req *ReservationRequest) (ReservationResponse, error) {
	log.Println("In Reservation CheckAvailability")
	
	res := ReservationResponse{
		HotelID: []string {"mus_1", "mus_2", "mus_3"},
	}

	return res, nil
}

