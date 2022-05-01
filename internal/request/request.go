package request

import (
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

type SearchParameters struct {
	Query   string
	Page    int
	PerPage int
}

// ParseSearchParameters returns correctly formatted search parameters for a
// GET request query parameters
func ParseSearchParameters(request events.APIGatewayProxyRequest) (SearchParameters, error) {
	query := request.QueryStringParameters["query"]
	if query == "" {
		return SearchParameters{}, errors.New("query is invalid")
	}

	page, err := strconv.Atoi(request.QueryStringParameters["page"])
	if err != nil {
		return SearchParameters{}, err
	}
	if page < 0 {
		return SearchParameters{}, errors.New("page parameter in negative")
	}

	perPage, err := strconv.Atoi(request.QueryStringParameters["perPage"])
	if err != nil {
		return SearchParameters{}, err
	}
	if perPage < 0 {
		return SearchParameters{}, errors.New("perPage parameter in negative")
	}

	parameters := SearchParameters{
		Query:   query,
		Page:    page,
		PerPage: perPage,
	}

	return parameters, nil
}
