# Go SDK for Azure OpenAI

This package provides Go Client
for [Azure OpenAI](https://azure.microsoft.com/en-us/products/cognitive-services/openai-service/)

## Install

Just type this script on your project.

```shell
$ go get github.com/piroyoung/go-aoai
```

## Example Usage

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	aoai "github.com/piroyoung/go-aoai"

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

```

Then we got belows.

```json
{
        "id": "cmpl-6zN8ZDK3V6fxT4jyKYGCPwyhroWKK",
        "object": "text_completion",
        "created": 1680084967,
        "model": "gpt-35-turbo",
        "choices": [
                {
                        "text": " the red hills of Georgia, the sons of former slaves and the sons of former slave owners will be able to sit together at the table of brotherhood.\n\nI have a dream today.\n\nI have a dream that one day, even the state of Mississippi, a state sweltering with the heat of injustice, sweltering with the heat of oppression, will be transformed into an oasis of freedom and justice.\n\nI have a dream today.\n\nI have a dream that my four little children will one day",
                        "logprobs": {},
                        "finish_reason": "length"
                }
        ]
}
```

## APIs

### Completion
```go
func (c *AzureOpenAI) Completion(ctx context.Context, request CompletionRequest) (*CompletionResponse, error)
```

Simple text completion api. If you pass some incomplete text, it will guess the rest of the text.

#### Usecase 
```go
request := CompletionRequest{
    Prompts:   []string{"I have a dream that one day on"},
    MaxTokens: 100,
    Stream:    false,
}

response, err := client.Completion(ctx, request)
```

### CompletionStream
```go
func (a *AzureOpenAI) CompletionStream(ctx context.Context, request CompletionRequest, consumer func(CompletionResponse) error) error
```
`CompletionStream` is a streaming api of `Completion`
If a field `CompletionRequest.Stream` is `true`, it will return a stream of responses with response header `Content-Type: text/event-stream'`.
We can process each chunk of response with `consumer` function.

#### Usecase
```go
request := CompletionRequest{
    Prompts:   []string{"I have a dream that one day on"},
    MaxTokens: 100,
    Stream:    true,
}

response, err := client.CompletionStream(ctx, request, func(chunk CompletionResponse) error {
    fmt.Println(chunk)
    return nil
})
```

### Embedding
```go
func (c *AzureOpenAI) Embedding(ctx context.Context, request EmbeddingRequest) (*EmbeddingResponse, error)
```
`Embedding` api returns a vector representation of the input text, and the vector is used to calculate the similarity between texts.

#### Usecase
```go
request := EmbeddingRequest{
    Prompts: []string{"I have a dream that one day on"},
}

response, err := client.Embedding(ctx, request)
```


### ChatCompletion
```go
func (a *AzureOpenAI) ChatCompletion(ctx context.Context, request ChatRequest) (*ChatResponse, error)
```
`ChatCompletion` api is a chatbot api. It can be used to generate a response to a given prompt.

#### Usecase
```go
request := ChatRequest{
	Messages: []ChatMessage{
		{
			Role:    "user",
			Content: "What is Azure OpenAI?",
		},
	},
	MaxTokens: 100,
}

response, err := client.ChatCompletion(ctx, request)
```

### ChatCompletionStream
```go
func (a *AzureOpenAI) ChatCompletionStream(ctx context.Context, request ChatRequest, consumer func(ChatResponse) error) error
```
`ChatCompletionStream` is a streaming api of `ChatCompletion`
If a field `ChatRequest.Stream` is `true`, it will return a stream of responses with response header `Content-Type: text/event-stream'`.
We can process each chunk of response with `consumer` function as same as `CompletionStream`.

#### Usecase
```go
request := ChatRequest{
	Messages: []ChatMessage{
		{
			Role:    "user",
			Content: "What is Azure OpenAI?",
		},
	},
	MaxTokens: 100,
	Stream:    true,
}

response, err := client.ChatCompletionStream(ctx, request, func(chunk ChatResponse) error {
    fmt.Println(chunk)
    return nil
})
```


## Global Parameters

This SDK requires some parameters to identify your project and deployment.

| name | description |
| :--- | :--- |
| `resourceName` | Name of Azure OpenAI resource |
| `deploymentName` | Name of deployment in your Azure OpenAI resource |
| `apiVersion` | Check [here](https://learn.microsoft.com/en-us/azure/cognitive-services/openai/reference).|
| `accessToken` | Authentication key. We can use Azure Active Directory Authentication(TBD). |

### `resourceName`

<img width="900" alt="resource_name" src="https://user-images.githubusercontent.com/6128022/228507736-f65f4a65-f1f2-4e34-b51c-22ef8e172948.png">

### `deploymentName`

<img width="900" alt="deployment_name" src="https://user-images.githubusercontent.com/6128022/228508010-3b65aced-c7b6-4ba9-b5e5-9b9b93e22b58.png">

### `accessToken`

<img width="900" alt="api_key" src="https://user-images.githubusercontent.com/6128022/228511558-3b42cf21-b5db-445a-9bfc-a672aac8a6f1.png">

## Request Parameters.

Models of Request/Response body are defined in `model.go`,
check [here](https://github.com/piroyoung/go-aoai/blob/main/model.go).
And you can also
see [Swagger API Reference](https://github.com/Azure/azure-rest-api-specs/blob/main/specification/cognitiveservices/data-plane/AzureOpenAI/inference/stable/2022-12-01/inference.json).
