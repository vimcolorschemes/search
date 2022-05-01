package request

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestParseSearchParameters(t *testing.T) {
	t.Run("should return all parameters and convert number parameters from string", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "page": "1", "perPage": "10", "backgrounds": "dark,light"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		result, err := ParseSearchParameters(request)
		if err != nil {
			t.Errorf("Incorrect result for ParseSearchParameters, got error: %s", err)
		}

		if result.Query != "test" {
			t.Errorf("Incorrect result for ParseSearchParameters, got query: %s, want query: %s", result.Query, "test")
		}

		if result.Page != 1 {
			t.Errorf("Incorrect result for ParseSearchParameters, got page: %d, want page: %d", result.Page, 1)
		}

		if result.PerPage != 10 {
			t.Errorf("Incorrect result for ParseSearchParameters, got perPage: %d, want perPage: %d", result.PerPage, 10)
		}

		if len(result.Backgrounds) != 2 {
			t.Errorf("Incorrect result for ParseSearchParameters, got backgrounds: %s, want backgrounds: %s", result.Backgrounds, []string{"dark", "light"})
		}
	})

	t.Run("should return error when the query is missing", func(t *testing.T) {
		parameters := map[string]string{"page": "1", "perPage": "10"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for query parameter missing")
		}
	})

	t.Run("should return error when the query is empty", func(t *testing.T) {
		parameters := map[string]string{"query": "", "page": "1", "perPage": "10"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for query parameter missing")
		}
	})

	t.Run("should return error when the page is missing", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "perPage": "10"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for page parameter missing")
		}
	})

	t.Run("should return error when the page is not a number", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "page": "test", "perPage": "10"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for page parameter not being a number")
		}
	})

	t.Run("should return error when the page is less than zero", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "page": "-10", "perPage": "10"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for page parameter being negative")
		}
	})

	t.Run("should return error when the perPage is missing", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "page": "10"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for perPage parameter missing")
		}
	})

	t.Run("should return error when the perPage is not a number", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "page": "10", "perPage": "test"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for perPage parameter not being a number")
		}
	})

	t.Run("should return error when the perPage is less than zero", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "page": "10", "perPage": "-10"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for perPage parameter being negative")
		}
	})

	t.Run("should return error when the backgrounds is missing", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "page": "10", "perPage": "10"}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for backgrounds missing")
		}
	})

	t.Run("should return error when the backgrounds is empty", func(t *testing.T) {
		parameters := map[string]string{"query": "test", "page": "10", "perPage": "10", "backgrounds": ""}
		request := events.APIGatewayProxyRequest{QueryStringParameters: parameters}

		_, err := ParseSearchParameters(request)
		if err == nil {
			t.Error("Incorrect result for ParseSearchParameters, expected error for backgrounds being empty")
		}
	})
}
