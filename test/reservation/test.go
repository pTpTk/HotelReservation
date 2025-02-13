package main

import (
	"context"
	"fmt"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

// The go equivalent of enum
type RequestType int

const (
	Reserve RequestType = iota
	Check
)

type ReservationRequest struct {
	RequestType  RequestType
	CustomerName string     
	HotelID    []string     
	InDate       string     
	OutDate      string     
	RoomNumber   int        
}


// main uses the AWS SDK for Go (v2) to create an AWS Lambda client and list up to 10
// functions in your account.
// This example uses the default settings specified in your shared credentials
// and config files.
func main() {
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	lambdaClient := lambda.NewFromConfig(sdkConfig)

	req := ReservationRequest{
		RequestType: Check,
		CustomerName: "mike",
		HotelID: []string{"id1", "id2"},
		InDate: "2002-12-12",
		OutDate: "2002-12-14",
		RoomNumber: 202,
	}

	payload, err := json.Marshal(req)

	output, err := lambdaClient.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String("reservation"),
		LogType:      types.LogTypeNone,
		Payload:      payload,
	})
	fmt.Printf("%s", output.Payload)
}
