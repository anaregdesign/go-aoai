package main

import (
	"fmt"
	"goaoai"
	"os"
)

func main() {
	resourceName := "resourceName"
	deploymentName := "deploymentName"
	apiVersion := "apiVersion"
	accessToken := os.Getenv("AZURE_OPENAI_API_KEY")

	client := goaoai.New(resourceName, deploymentName, apiVersion, accessToken)

	result, _ := client.Completion("Negate the following sentence.The price for bubblegum increased on thursday.\n\n Negated Sentence:", 50)
	fmt.Println(result)
}
