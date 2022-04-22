package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vimcolorschemes/search/internal/database"
)

type Event struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

func HandleLambdaEvent(event Event) (interface{}, error) {
	searchIndex := database.GetSearchIndex()
	return searchIndex, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
