package request

import (
	"encoding/json"
	"errors"
	"strings"

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
	if page < 1 {
		return SearchParameters{}, errors.New("page parameter in smaller than 1")
	}

	perPage, err := strconv.Atoi(request.QueryStringParameters["perPage"])
	if err != nil {
		return SearchParameters{}, err
	}
	if perPage < 1 {
		return SearchParameters{}, errors.New("perPage parameter smaller than 1")
	}

	parameters := SearchParameters{
		Query:   query,
		Page:    page,
		PerPage: perPage,
	}

	return parameters, nil
}

// BuildErrorBody builds a JSON request response body using an error message
func BuildErrorBody(messages ...string) string {
	body := make(map[string]string)
	body["message"] = strings.Join(messages, " ")

	payload, err := json.Marshal(body)
	if err != nil {
		return ""
	}

	return string(payload)
}
