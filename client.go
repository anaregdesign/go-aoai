package aoai

import (
	m "aoai/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AzureOpenAI struct {
	client             *http.Client
	resourceName       string
	deploymentName     string
	apiVersion         string
	useActiveDirectory bool
	accessToken        string
}

func NewWithActiveDirectory(resourceName string, deploymentName string, apiVersion string, accessToken string) *AzureOpenAI {
	return &AzureOpenAI{
		client:             &http.Client{},
		resourceName:       resourceName,
		deploymentName:     deploymentName,
		apiVersion:         apiVersion,
		useActiveDirectory: true,
		accessToken:        accessToken,
	}
}

func New(resourceName string, deploymentName string, apiVersion string, accessToken string) *AzureOpenAI {
	return &AzureOpenAI{
		client:             &http.Client{},
		resourceName:       resourceName,
		deploymentName:     deploymentName,
		apiVersion:         apiVersion,
		useActiveDirectory: false,
		accessToken:        accessToken,
	}
}

func (a *AzureOpenAI) endpoint() string {
	return fmt.Sprintf("https://%s.openai.azure.com/openai/deployments/%s", a.resourceName, a.deploymentName)
}

func (a *AzureOpenAI) header() http.Header {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	if a.useActiveDirectory {
		header.Add("Authorization", fmt.Sprintf("Bearer %s", a.accessToken))
	} else {
		header.Add("api-key", a.accessToken)
	}
	return header
}

func (a *AzureOpenAI) Completion(prompt []string, maxTokens int) (*m.CompletionResponse, error) {
	endpoint := fmt.Sprintf("%s/completions?api-version=%s", a.endpoint(), a.apiVersion)

	completionRequest := m.CompletionRequest{
		Prompt:    prompt,
		MaxTokens: maxTokens,
	}
	requestBody, _ := json.Marshal(completionRequest)
	request, _ := http.NewRequest("POST", endpoint, bytes.NewReader(requestBody))
	request.Header = a.header()

	response, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		var errorResponse m.ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return nil, err
		}
		return nil, &errorResponse.Error
	} else {
		var completionResponse m.CompletionResponse
		if err := json.Unmarshal(responseBody, &completionResponse); err != nil {
			return nil, err
		}
		return &completionResponse, nil
	}
}
