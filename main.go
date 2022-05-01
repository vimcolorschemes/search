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
		return store(request), nil
	case "GET":
		return search(request), nil
	default:
		body := req.BuildErrorBody("HTTP method not supported")
		return events.APIGatewayProxyResponse{Body: body, StatusCode: 400, Headers: Headers}, nil
	}
}

func store(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var searchIndex []repository.Repository

	if err := json.Unmarshal([]byte(request.Body), &searchIndex); err != nil {
		body := req.BuildErrorBody("error trying to parse request body:", err.Error())
		return events.APIGatewayProxyResponse{Body: body, StatusCode: 500, Headers: Headers}
	}

	if err := database.Store(searchIndex); err != nil {
		body := req.BuildErrorBody("error trying to store the search index:", err.Error())
		return events.APIGatewayProxyResponse{Body: body, StatusCode: 500, Headers: Headers}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Headers: Headers}
}

func search(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	parameters, err := req.ParseSearchParameters(request)
	if err != nil {
		body := req.BuildErrorBody("error parsing search parameters:", err.Error())
		return events.APIGatewayProxyResponse{Body: body, StatusCode: 400, Headers: Headers}
	}

	repositories, total, err := database.Search(parameters)
	if err != nil {
		body := req.BuildErrorBody("error storing search index:", err.Error())
		return events.APIGatewayProxyResponse{Body: body, StatusCode: 500, Headers: Headers}
	}

	result := map[string]interface{}{"repositories": repositories, "totalCount": total}

	payload, err := json.Marshal(result)
	if err != nil {
		body := req.BuildErrorBody("error encoding search result to JSON:", err.Error())
		return events.APIGatewayProxyResponse{Body: body, StatusCode: 500, Headers: Headers}
	}

	return events.APIGatewayProxyResponse{Body: string(payload), StatusCode: 200, Headers: Headers}
}

func main() {
	lambda.Start(handle)
}
