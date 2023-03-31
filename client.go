package aoai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type AzureOpenAI struct {
	httpClient         *http.Client
	resourceName       string
	deploymentName     string
	apiVersion         string
	useActiveDirectory bool
	accessToken        string
}

func NewWithActiveDirectory(resourceName string, deploymentName string, apiVersion string, accessToken string) *AzureOpenAI {
	return &AzureOpenAI{
		httpClient:         &http.Client{},
		resourceName:       resourceName,
		deploymentName:     deploymentName,
		apiVersion:         apiVersion,
		useActiveDirectory: true,
		accessToken:        accessToken,
	}
}

func New(resourceName string, deploymentName string, apiVersion string, accessToken string) *AzureOpenAI {
	return &AzureOpenAI{
		httpClient:         &http.Client{},
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

func (a *AzureOpenAI) Completion(ctx context.Context, completionRequest CompletionRequest) (*CompletionResponse, error) {
	if completionRequest.Stream {
		return nil, fmt.Errorf("streaming is not supported. Try `CompletionStream` instead")
	}

	endpoint := fmt.Sprintf("%s/completions?api-version=%s", a.endpoint(), a.apiVersion)

	requestBody, _ := json.Marshal(completionRequest)
	request, _ := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(requestBody))
	request.Header = a.header()

	response, err := a.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return nil, err
		}
		return nil, &errorResponse.Error
	} else {
		var completionResponse CompletionResponse
		if err := json.Unmarshal(responseBody, &completionResponse); err != nil {
			return nil, err
		}
		return &completionResponse, nil
	}
}

func (a *AzureOpenAI) Embedding(ctx context.Context, embeddingRequest EmbeddingRequest) (*EmbeddingResponse, error) {
	endpoint := fmt.Sprintf("%s/embeddings?api-version=%s", a.endpoint(), a.apiVersion)

	requestBody, _ := json.Marshal(embeddingRequest)
	request, _ := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(requestBody))
	request.Header = a.header()

	response, err := a.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return nil, err
		}
		return nil, &errorResponse.Error
	} else {
		var embeddingResponse EmbeddingResponse
		if err := json.Unmarshal(responseBody, &embeddingResponse); err != nil {
			return nil, err
		}
		return &embeddingResponse, nil
	}
}

func (a *AzureOpenAI) ChatCompletion(ctx context.Context, chatRequest ChatRequest) (*ChatResponse, error) {
	if chatRequest.Stream {
		return nil, fmt.Errorf("streaming is not supported. Try `ChatCompletionStream` instead")
	}

	endpoint := fmt.Sprintf("%s/chat/completions?api-version=%s", a.endpoint(), a.apiVersion)

	requestBody, _ := json.Marshal(chatRequest)
	request, _ := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(requestBody))
	request.Header = a.header()

	response, err := a.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return nil, err
		}
		return nil, &errorResponse.Error
	} else {
		var chatResponse ChatResponse
		if err := json.Unmarshal(responseBody, &chatResponse); err != nil {
			return nil, err
		}
		return &chatResponse, nil
	}
}

// CompletionStream
// https://learn.microsoft.com/en-us/azure/cognitive-services/openai/reference
// Whether to stream back partial progress. If set, tokens will be sent as data-only server-sent events as they become
// available, with the stream terminated by a `data: [DONE]` message.
func (a *AzureOpenAI) CompletionStream(ctx context.Context, completionRequest CompletionRequest, consumer func(completionResponse CompletionResponse) error) error {
	if !completionRequest.Stream {
		return fmt.Errorf("streaming is not enabled. Try `Completion` instead")
	}

	endpoint := fmt.Sprintf("%s/completions?api-version=%s", a.endpoint(), a.apiVersion)

	requestBody, _ := json.Marshal(completionRequest)
	request, _ := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(requestBody))
	request.Header = a.header()

	response, err := a.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		responseBody, _ := io.ReadAll(response.Body)
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return err
		}
		return &errorResponse.Error
	}

	reader := bufio.NewReader(response.Body)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			m, err := reader.ReadString('\n')
			if err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}

			// remove prefix 'data: ' and suffix '\n'
			m = strings.TrimPrefix(m, "data: ")
			m = strings.TrimSuffix(m, "\n")
			if m == "" {
				// stream is delimited by '\n\n'
				continue
			} else if m == "[DONE]" {
				// stream is terminated by a `data: [DONE]` message
				return nil
			}

			var completionResponse CompletionResponse
			if err := json.Unmarshal([]byte(m), &completionResponse); err != nil {
				return err
			}
			if err := consumer(completionResponse); err != nil {
				return err
			}
		}
	}
}

// ChatCompletionStream
// https://learn.microsoft.com/en-us/azure/cognitive-services/openai/reference
// Whether to stream back partial progress. If set, tokens will be sent as data-only server-sent events as they become
// available, with the stream terminated by a `data: [DONE]` message.
func (a *AzureOpenAI) ChatCompletionStream(ctx context.Context, chatRequest ChatRequest, consumer func(chatResponse ChatResponse) error) error {
	if !chatRequest.Stream {
		return fmt.Errorf("streaming is not enabled. Try `ChatCompletion` instead")
	}
	endpoint := fmt.Sprintf("%s/chat/completions?api-version=%s", a.endpoint(), a.apiVersion)

	requestBody, _ := json.Marshal(chatRequest)
	request, _ := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(requestBody))
	request.Header = a.header()

	response, err := a.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		responseBody, _ := io.ReadAll(response.Body)
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return err
		}
		return &errorResponse.Error
	}

	reader := bufio.NewReader(response.Body)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			m, err := reader.ReadString('\n')
			if err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}

			// remove prefix 'data: ' and suffix '\n'
			m = strings.TrimPrefix(m, "data: ")
			m = strings.TrimSuffix(m, "\n")
			if m == "" {
				// stream is delimited by '\n\n'
				continue
			} else if m == "[DONE]" {
				// stream is terminated by a `data: [DONE]` message
				return nil
			}

			var chatResponse ChatResponse
			if err := json.Unmarshal([]byte(m), &chatResponse); err != nil {
				return err
			}
			if err := consumer(chatResponse); err != nil {
				return err
			}
		}
	}
}
