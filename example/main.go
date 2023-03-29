package main

import (
	"encoding/json"
	"fmt"
	"goaoai"
	"os"
)

func main() {
	resourceName := "example-aoai-02"
	deploymentName := "gpt-35-turbo-0301"
	apiVersion := "2023-03-15-preview"
	accessToken := os.Getenv("AZURE_OPENAI_API_KEY")

	client := goaoai.New(resourceName, deploymentName, apiVersion, accessToken)

	result, _ := client.Completion("I have a dream that one day on", 50)
	if jsonString, err := json.MarshalIndent(result, "", "\t"); err == nil {
		fmt.Println(string(jsonString))
	}
}
