package aoai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestAzureOpenAI_ChatCompletion(t *testing.T) {
	type fields struct {
		httpClient         *http.Client
		resourceName       string
		deploymentName     string
		apiVersion         string
		useActiveDirectory bool
		accessToken        string
	}
	type args struct {
		ctx         context.Context
		chatRequest ChatRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ChatResponse
		wantErr bool
	}{
		{
			name: "validCase",
			fields: fields{
				httpClient:         &http.Client{},
				resourceName:       "example-aoai-02",
				deploymentName:     "gpt-35-turbo-0301",
				apiVersion:         "2023-03-15-preview",
				useActiveDirectory: false,
				accessToken:        os.Getenv("AZURE_OPENAI_API_KEY"),
			},
			args: args{
				ctx: context.Background(),
				chatRequest: ChatRequest{
					Messages: []ChatMessage{
						{
							Role:    "user",
							Content: "What is Azure OpenAI?",
						},
					},
					MaxTokens: 100,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AzureOpenAI{
				httpClient:         tt.fields.httpClient,
				resourceName:       tt.fields.resourceName,
				deploymentName:     tt.fields.deploymentName,
				apiVersion:         tt.fields.apiVersion,
				useActiveDirectory: tt.fields.useActiveDirectory,
				accessToken:        tt.fields.accessToken,
			}
			got, err := a.ChatCompletion(tt.args.ctx, tt.args.chatRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if jsonString, _ := json.MarshalIndent(got, "", "\t"); jsonString != nil {
				fmt.Println(string(jsonString))
			}
		})
	}
}

func TestAzureOpenAI_Completion(t *testing.T) {
	type fields struct {
		httpClient         *http.Client
		resourceName       string
		deploymentName     string
		apiVersion         string
		useActiveDirectory bool
		accessToken        string
	}
	type args struct {
		ctx               context.Context
		completionRequest CompletionRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CompletionResponse
		wantErr bool
	}{
		{
			name: "validCase",
			fields: fields{
				httpClient:         &http.Client{},
				resourceName:       "example-aoai-02",
				deploymentName:     "gpt-35-turbo-0301",
				apiVersion:         "2023-03-15-preview",
				useActiveDirectory: false,
				accessToken:        os.Getenv("AZURE_OPENAI_API_KEY"),
			},
			args: args{
				ctx: context.Background(),
				completionRequest: CompletionRequest{
					Prompts:   []string{"I have a dream that one day on"},
					MaxTokens: 100,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AzureOpenAI{
				httpClient:         tt.fields.httpClient,
				resourceName:       tt.fields.resourceName,
				deploymentName:     tt.fields.deploymentName,
				apiVersion:         tt.fields.apiVersion,
				useActiveDirectory: tt.fields.useActiveDirectory,
				accessToken:        tt.fields.accessToken,
			}
			got, err := a.Completion(tt.args.ctx, tt.args.completionRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Completion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if jsonString, _ := json.MarshalIndent(got, "", "\t"); jsonString != nil {
				fmt.Println(string(jsonString))
			}
		})
	}
}

func TestAzureOpenAI_Embedding(t *testing.T) {
	type fields struct {
		httpClient         *http.Client
		resourceName       string
		deploymentName     string
		apiVersion         string
		useActiveDirectory bool
		accessToken        string
	}
	type args struct {
		ctx              context.Context
		embeddingRequest EmbeddingRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *EmbeddingResponse
		wantErr bool
	}{
		{
			name: "validCase",
			fields: fields{
				httpClient:         &http.Client{},
				resourceName:       "example-aoai-02",
				deploymentName:     "text-embedding-ada-002",
				apiVersion:         "2023-03-15-preview",
				useActiveDirectory: false,
				accessToken:        os.Getenv("AZURE_OPENAI_API_KEY"),
			},
			args: args{
				ctx: context.Background(),
				embeddingRequest: EmbeddingRequest{
					Inputs: []string{"I love both Microsoft and OpenSource."},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AzureOpenAI{
				httpClient:         tt.fields.httpClient,
				resourceName:       tt.fields.resourceName,
				deploymentName:     tt.fields.deploymentName,
				apiVersion:         tt.fields.apiVersion,
				useActiveDirectory: tt.fields.useActiveDirectory,
				accessToken:        tt.fields.accessToken,
			}
			got, err := a.Embedding(tt.args.ctx, tt.args.embeddingRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Embedding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if jsonString, _ := json.MarshalIndent(got, "", "\t"); jsonString != nil {
				fmt.Println(string(jsonString))
			}
		})
	}
}

func TestAzureOpenAI_endpoint(t *testing.T) {
	type fields struct {
		httpClient         *http.Client
		resourceName       string
		deploymentName     string
		apiVersion         string
		useActiveDirectory bool
		accessToken        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "validCase",
			fields: fields{
				httpClient:         &http.Client{},
				resourceName:       "example-aoai-02",
				deploymentName:     "gpt-35-turbo-0301",
				apiVersion:         "2023-03-15-preview",
				useActiveDirectory: false,
				accessToken:        "dummy",
			},
			want: "https://example-aoai-02.openai.azure.com/openai/deployments/gpt-35-turbo-0301",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AzureOpenAI{
				httpClient:         tt.fields.httpClient,
				resourceName:       tt.fields.resourceName,
				deploymentName:     tt.fields.deploymentName,
				apiVersion:         tt.fields.apiVersion,
				useActiveDirectory: tt.fields.useActiveDirectory,
				accessToken:        tt.fields.accessToken,
			}
			if got := a.endpoint(); got != tt.want {
				t.Errorf("endpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAzureOpenAI_header(t *testing.T) {
	want := http.Header{}
	want.Add("Content-Type", "application/json")
	want.Add("api-key", "some API key")

	type fields struct {
		httpClient         *http.Client
		resourceName       string
		deploymentName     string
		apiVersion         string
		useActiveDirectory bool
		accessToken        string
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Header
	}{
		{
			name: "withAPIKey",
			fields: fields{
				httpClient:         &http.Client{},
				resourceName:       "example-aoai-02",
				deploymentName:     "gpt-35-turbo-0301",
				apiVersion:         "2023-03-15-preview",
				useActiveDirectory: false,
				accessToken:        "some API key",
			},
			want: want,
		},
		{
			name: "withAccessToken",
			fields: fields{
				httpClient:         &http.Client{},
				resourceName:       "example-aoai-02",
				deploymentName:     "gpt-35-turbo-0301",
				apiVersion:         "2023-03-15-preview",
				useActiveDirectory: true,
				accessToken:        "some access token",
			},
			want: http.Header{
				"Content-Type":  []string{"application/json"},
				"Authorization": []string{"Bearer some access token"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AzureOpenAI{
				httpClient:         tt.fields.httpClient,
				resourceName:       tt.fields.resourceName,
				deploymentName:     tt.fields.deploymentName,
				apiVersion:         tt.fields.apiVersion,
				useActiveDirectory: tt.fields.useActiveDirectory,
				accessToken:        tt.fields.accessToken,
			}
			if got := a.header(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("header() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAzureOpenAI_CompletionStream(t *testing.T) {
	resourceName := "example-aoai-02"
	deploymentName := "gpt-35-turbo-0301"
	apiVersion := "2023-03-15-preview"
	accessToken := os.Getenv("AZURE_OPENAI_API_KEY")

	client := New(resourceName, deploymentName, apiVersion, accessToken)
	ctx := context.Background()
	request := CompletionRequest{
		Prompts:   []string{"I love both Microsoft and OpenSource, because "},
		MaxTokens: 100,
		Stream:    true,
	}
	err := client.CompletionStream(ctx, request, func(res CompletionResponse) error {
		fmt.Print(res.Choices[0].Text)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
