package main

import (
	"log"
	"strconv"
	"context"
	"fmt"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	lambdaSDK"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func main() {
	log.Println("Starting service...")
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	path := req.RawPath
	if path == "/recommendations" {
		return recommendHandler(ctx, req)
	}

	return events.APIGatewayV2HTTPResponse{StatusCode: 300, Body: ""}, nil
}

type RecommendationRequest struct {
	Require string
	Lat     float64 
	Lon     float64 
}

func recommendHandler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	sLat := req.QueryStringParameters["lat"]
	sLon := req.QueryStringParameters["lon"]
	
	Lat, _ := strconv.ParseFloat(sLat, 64)
	lat := float64(Lat)
	Lon, _ := strconv.ParseFloat(sLon, 64)
	lon := float64(Lon)
	
	var lambdaReq RecommendationRequest
	lambdaReq.Require = req.QueryStringParameters["require"]
	lambdaReq.Lat = lat
	lambdaReq.Lon = lon

	fmt.Print(lambdaReq, "\n")

	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: 300, Body: ""}, err
	}
	lambdaClient := lambdaSDK.NewFromConfig(sdkConfig)

	payload, err := json.Marshal(req)

	output, err := lambdaClient.Invoke(ctx, &lambdaSDK.InvokeInput{
		FunctionName: aws.String("recommendation"),
		LogType:      types.LogTypeNone,
		Payload:      payload,
	})
	fmt.Printf("%s", output.Payload)

	return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: "hello"}, nil
}