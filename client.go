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
	return postJsonRequest[CompletionRequest, CompletionResponse](ctx, a.httpClient, endpoint, a.header(), completionRequest)

}

func (a *AzureOpenAI) Embedding(ctx context.Context, embeddingRequest EmbeddingRequest) (*EmbeddingResponse, error) {
	endpoint := fmt.Sprintf("%s/embeddings?api-version=%s", a.endpoint(), a.apiVersion)
	return postJsonRequest[EmbeddingRequest, EmbeddingResponse](ctx, a.httpClient, endpoint, a.header(), embeddingRequest)
}

func (a *AzureOpenAI) ChatCompletion(ctx context.Context, chatRequest ChatRequest) (*ChatResponse, error) {
	if chatRequest.Stream {
		return nil, fmt.Errorf("streaming is not supported. Try `ChatCompletionStream` instead")
	}

	endpoint := fmt.Sprintf("%s/chat/completions?api-version=%s", a.endpoint(), a.apiVersion)
	return postJsonRequest[ChatRequest, ChatResponse](ctx, a.httpClient, endpoint, a.header(), chatRequest)
}

func (a *AzureOpenAI) CompletionStream(ctx context.Context, completionRequest CompletionRequest, consumer func(completionResponse CompletionResponse) error) error {
	if !completionRequest.Stream {
		return fmt.Errorf("streaming is not enabled. Try `Completion` instead")
	}

	endpoint := fmt.Sprintf("%s/completions?api-version=%s", a.endpoint(), a.apiVersion)
	return postJsonRequestStream[CompletionRequest, CompletionResponse](ctx, a.httpClient, endpoint, a.header(), completionRequest, consumer)
}

// ChatCompletionStream

func (a *AzureOpenAI) ChatCompletionStream(ctx context.Context, chatRequest ChatRequest, consumer func(chatResponse ChatResponse) error) error {
	if !chatRequest.Stream {
		return fmt.Errorf("streaming is not enabled. Try `ChatCompletion` instead")
	}
	endpoint := fmt.Sprintf("%s/chat/completions?api-version=%s", a.endpoint(), a.apiVersion)
	return postJsonRequestStream[ChatRequest, ChatResponse](ctx, a.httpClient, endpoint, a.header(), chatRequest, consumer)
}

func postJsonRequest[S, T any](ctx context.Context, httpClient *http.Client, endpoint string, header http.Header, request S) (*T, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	httpRequest, _ := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(requestBody))
	httpRequest.Header = header

	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	responseBody, _ := io.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode != 200 {
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return nil, err
		}
		return nil, &errorResponse.Error
	} else {
		var response T
		if err := json.Unmarshal(responseBody, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}
}

// postJsonRequestStream
// https://learn.microsoft.com/en-us/azure/cognitive-services/openai/reference
// Whether to stream back partial progress. If set, tokens will be sent as data-only server-sent events as they become
// available, with the stream terminated by a `data: [DONE]` message.
func postJsonRequestStream[S, T any](ctx context.Context, httpClient *http.Client, endpoint string, header http.Header, request S, consumer func(response T) error) error {

	requestBody, _ := json.Marshal(request)
	httpRequest, _ := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(requestBody))
	httpRequest.Header = header

	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		responseBody, _ := io.ReadAll(httpResponse.Body)
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return err
		}
		return &errorResponse.Error
	}

	reader := bufio.NewReader(httpResponse.Body)

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

			var chunk T
			if err := json.Unmarshal([]byte(m), &chunk); err != nil {
				return err
			}
			if err := consumer(chunk); err != nil {
				return err
			}
		}
	}
}
