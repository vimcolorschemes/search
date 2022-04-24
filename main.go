package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vimcolorschemes/search/internal/database"
)

const (
	Store  = "store"
	Search = "search"
)

var Headers = map[string]string{"Access-Control-Allow-Origin": "*"}

func handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return store(request.Body)
	case "GET":
		return search(request.QueryStringParameters["term"])
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: Headers}, nil
	}
}

func store(payload string) (events.APIGatewayProxyResponse, error) {
	var searchIndex interface{}

	if err := json.Unmarshal([]byte(payload), &searchIndex); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: Headers}, err
	}

	if err := database.StoreSearchIndex(searchIndex); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: Headers}, err
	}

	return events.APIGatewayProxyResponse{Body: "success", StatusCode: 200, Headers: Headers}, nil
}

func search(term string) (events.APIGatewayProxyResponse, error) {
	searchIndex := database.GetSearchIndex()
	payload, _ := json.Marshal(searchIndex)
	return events.APIGatewayProxyResponse{Body: string(payload), StatusCode: 200, Headers: Headers}, nil
}

func main() {
	lambda.Start(handle)
}
