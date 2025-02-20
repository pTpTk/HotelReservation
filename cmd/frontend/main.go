package main

import (
	"log"
	"strconv"
	"context"
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	lambdaSDK"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

	"github.com/pTpTk/HotelReservation/proto"
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

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusBadRequest, 
		Body: "unsupported request"
	}, nil
}

type RecommendationRequest struct {
	Require string
	Lat     float64 
	Lon     float64 
}

type RecommendationResponse struct {
	HotelIDs []string
}

func recommendHandler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// initialize lambda client
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: 300, Body: ""}, err
	}
	lambdaClient := lambdaSDK.NewFromConfig(sdkConfig)

	// get query parameters
	sLat := req.QueryStringParameters["lat"]
	sLon := req.QueryStringParameters["lon"]
	if sLat == "" || sLon == "" {
		panic("Please specify location params")
	}
	
	Lat, _ := strconv.ParseFloat(sLat, 64)
	lat := float64(Lat)
	Lon, _ := strconv.ParseFloat(sLon, 64)
	lon := float64(Lon)

	require := req.QueryStringParameters["require"]
	if require != "dis" && require != "rate" && require != "price" {
		panic("Please specify require params")
	}

	var lambdaReq RecommendationRequest
	lambdaReq.Require = req.QueryStringParameters["require"]
	lambdaReq.Lat = lat
	lambdaReq.Lon = lon

	payload, err := json.Marshal(req)

	// recommend hotels
	recResp, err := lambdaClient.Invoke(ctx, &lambdaSDK.InvokeInput{
		FunctionName: aws.String("recommendation"),
		LogType:      types.LogTypeNone,
		Payload:      payload,
	})
	if err != nil {
		fmt.Println("can't call recommendation")
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError, Body: ""}, nil
	}

	// grab locale from query params or default to en
	locale := req.QueryStringParameters["locale"]
	if locale == "" {
		locale = "en"
	}

	// hotel profiles
	profileResp, err := lambdaClient.Invoke(ctx, &lambdaSDK.InvokeInput{
		FunctionName: aws.String("profile"),
		LogType:      types.LogTypeNone,
		Payload:      payload,
	})
	if err != nil {
		fmt.Println("can't call recommendation")
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError, Body: ""}, nil
	}

	fmt.Printf("%s", output.Payload)

	return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: "hello"}, nil
}
