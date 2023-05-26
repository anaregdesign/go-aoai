package main

import (
	"context"
	"encoding/json"
	"fmt"
	aoai "github.com/anaregdesign/go-aoai"

	"os"
)

func main() {
	ctx := context.Background()

	resourceName := "example-aoai-02"
	deploymentName := "gpt-35-turbo-0301"
	apiVersion := "2023-03-15-preview"
	accessToken := os.Getenv("AZURE_OPENAI_API_KEY")

	client := aoai.New(resourceName, deploymentName, apiVersion, accessToken)

	request := aoai.CompletionRequest{
		Prompts:   []string{"I have a dream that one day on"},
		MaxTokens: 100,
		Stream:    false,
	}

	response, err := client.Completion(ctx, request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if jsonString, err := json.MarshalIndent(response, "", "\t"); err == nil {
		fmt.Println(string(jsonString))
	}
}
