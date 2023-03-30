package aoai

import (
	"context"
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
			fmt.Println(got)
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Completion() got = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Embedding() got = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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

func TestNew(t *testing.T) {
	type args struct {
		resourceName   string
		deploymentName string
		apiVersion     string
		accessToken    string
	}
	tests := []struct {
		name string
		args args
		want *AzureOpenAI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.resourceName, tt.args.deploymentName, tt.args.apiVersion, tt.args.accessToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithActiveDirectory(t *testing.T) {
	type args struct {
		resourceName   string
		deploymentName string
		apiVersion     string
		accessToken    string
	}
	tests := []struct {
		name string
		args args
		want *AzureOpenAI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithActiveDirectory(tt.args.resourceName, tt.args.deploymentName, tt.args.apiVersion, tt.args.accessToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithActiveDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
