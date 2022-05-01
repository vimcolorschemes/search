package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vimcolorschemes/search/internal/database"
	"github.com/vimcolorschemes/search/internal/repository"
	req "github.com/vimcolorschemes/search/internal/request"
)

const (
	Store  = "store"
	Search = "search"
)

var Headers = map[string]string{"Access-Control-Allow-Origin": "*"}

func handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return store(request)
	case "GET":
		return search(request)
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: Headers}, nil
	}
}

func store(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var searchIndex []repository.Repository

	if err := json.Unmarshal([]byte(request.Body), &searchIndex); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: Headers}, err
	}

	if err := database.Store(searchIndex); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: Headers}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Headers: Headers}, nil
}

func search(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	parameters, err := req.ParseSearchParameters(request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: Headers}, err
	}

	repositories, totalCount, err := database.Search(parameters)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: Headers}, err
	}

	result := map[string]interface{}{"repositories": repositories, "totalCount": totalCount}

	payload, err := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: Headers}, err
	}

	return events.APIGatewayProxyResponse{Body: string(payload), StatusCode: 200, Headers: Headers}, nil
}

func main() {
	lambda.Start(handle)
}
