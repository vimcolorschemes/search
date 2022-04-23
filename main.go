package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vimcolorschemes/search/internal/database"
)

const (
	Store  = "store"
	Search = "search"
)

type Event struct {
	Action  string `json:"action"`
	Payload string `json:"payload"`
}

func handle(event Event) (interface{}, error) {
	switch event.Action {
	case Store:
		return store(event.Payload)
	case Search:
		return search(event.Payload)
	default:
		return "", nil
	}
}

func store(payload string) (interface{}, error) {
	var searchIndex map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &searchIndex); err != nil {
		return "", err
	}

	if err := database.StoreSearchIndex(searchIndex); err != nil {
		return "", err
	}

	return "Success", nil
}

func search(term string) (interface{}, error) {
	searchIndex := database.GetSearchIndex()
	return searchIndex, nil
}

func main() {
	lambda.Start(handle)
}
