package main

import (
	"context"
	"log"
)

func GetRates(ctx context.Context, req *RateRequest) (RateResponse, error) {
	log.Println("In Rate GetRates")
	
	res := RateResponse{
		RatePlans: []RatePlan{
			RatePlan{
				HotelID: "hot_1",
				Code: "0300254",
				InDate: "2022-12-12",
				OutDate: "2022-12-14",
				RoomType: RoomType{
					BookableRate: 123,
					TotalRate: 345,
					TotalRateInclusive: 666,
					Code: "030024",
					Currency: "USD",
					RoomDescription: "not good",
				},
			},
		},
	}

	return res, nil
}

