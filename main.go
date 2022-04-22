package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vimcolorschemes/search/internal/database"
)

const (
	Store  = "store"
	Search = "search"
)

type Event struct {
	Action string `json:"action"`
}

func HandleLambdaEvent(event Event) (interface{}, error) {
	switch event.Action {
	case Store:
		return store()
	case Search:
		return search()
	default:
		return "", nil
	}
}

func store() (interface{}, error) {
	return "", nil
}

func search() (interface{}, error) {
	searchIndex := database.GetSearchIndex()
	return searchIndex, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
