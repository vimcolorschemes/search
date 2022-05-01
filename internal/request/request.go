package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

type SearchParameters struct {
	Query       string
	Page        int
	PerPage     int
	Backgrounds []string
}

const (
	Light = "light"
	Dark  = "dark"
)

// ParseSearchParameters returns correctly formatted search parameters for a
// GET request query parameters
func ParseSearchParameters(request events.APIGatewayProxyRequest) (SearchParameters, error) {
	query := request.QueryStringParameters["query"]
	if query == "" {
		return SearchParameters{}, errors.New("query is invalid")
	}

	page, err := strconv.Atoi(request.QueryStringParameters["page"])
	if err != nil {
		return SearchParameters{}, errors.New(fmt.Sprintf("error parsing '%s' as integer", request.QueryStringParameters["page"]))
	}
	if page < 1 {
		return SearchParameters{}, errors.New("page parameter is smaller than 1")
	}

	perPage, err := strconv.Atoi(request.QueryStringParameters["perPage"])
	if err != nil {
		return SearchParameters{}, errors.New(fmt.Sprintf("error parsing '%s' as integer", request.QueryStringParameters["perPage"]))
	}
	if perPage < 1 {
		return SearchParameters{}, errors.New("perPage parameter is smaller than 1")
	}

	backgrounds := strings.Split(request.QueryStringParameters["backgrounds"], ",")
	for _, background := range backgrounds {
		if background != Light && background != Dark {
			summary := fmt.Sprintf("['%s']", strings.Join(backgrounds, "', '"))
			return SearchParameters{}, errors.New(fmt.Sprintf("at least one background is invalid in %s. valid values are ['light', 'dark']", summary))
		}
	}

	parameters := SearchParameters{
		Query:       query,
		Page:        page,
		PerPage:     perPage,
		Backgrounds: backgrounds,
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
