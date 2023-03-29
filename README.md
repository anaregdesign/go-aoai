# Go Azure OpenAI
This package provides Go Client for [Azure OpenAI](https://azure.microsoft.com/en-us/products/cognitive-services/openai-service/)

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
Models of Request/Response body are defined in `model.go`, check [here](https://github.com/piroyoung/go-aoai/blob/main/model.go).
And you can also see [Swagger API Reference](https://github.com/Azure/azure-rest-api-specs/blob/main/specification/cognitiveservices/data-plane/AzureOpenAI/inference/stable/2022-12-01/inference.json).
