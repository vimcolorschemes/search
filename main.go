package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vimcolorschemes/search/internal/dotenv"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	test, exists := dotenv.Get("SEARCH_TEST")

	if exists {
		return MyResponse{Message: fmt.Sprintf("%s is %d years old! %s", event.Name, event.Age, test)}, nil
	}

	return MyResponse{Message: fmt.Sprintf("%s is %d years old! no dotenv", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
