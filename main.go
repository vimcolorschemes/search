package main

import (
	"encoding/json"
	"errors"
	"strconv"

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
		return search(request.QueryStringParameters)
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

func search(parameters map[string]string) (events.APIGatewayProxyResponse, error) {
	query, page, perPage, err := getSearchParameters(parameters)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: Headers}, err
	}

	repositories, totalCount := database.Search(query, page, perPage)

	result := map[string]interface{}{"repositories": repositories, "totalCount": totalCount}

	payload, err := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: Headers}, err
	}

	return events.APIGatewayProxyResponse{Body: string(payload), StatusCode: 200, Headers: Headers}, nil
}

func getSearchParameters(parameters map[string]string) (string, int, int, error) {
	query := parameters["query"]
	if query == "" {
		return "", -1, -1, errors.New("query is invalid")
	}

	page, err := strconv.Atoi(parameters["page"])
	if err != nil {
		return "", -1, -1, err
	}

	perPage, err := strconv.Atoi(parameters["perPage"])
	if err != nil {
		return "", -1, -1, err
	}

	return query, page, perPage, nil
}

func main() {
	lambda.Start(handle)
}
