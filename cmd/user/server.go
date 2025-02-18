package main

import (
	"context"
	"log"
)

func CheckUser(ctx context.Context, req *UserRequest) (UserResponse, error) {
	log.Println("In User CheckUser")
	
	res := UserResponse{
		Correct: true,
	}

	return res, nil
}

