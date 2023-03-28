package goaoai

// {
//  "prompt": "Negate the following sentence.The price for bubblegum increased on thursday.\n\n Negated Sentence:",
//  "max_tokens": 50
//}

type CompletionRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

// {
//  "model": "davinci",
//  "object": "text_completion",
//  "id": "cmpl-4509KAos68kxOqpE2uYGw81j6m7uo",
//  "created": 1637097562,
//  "choices": [
//    {
//      "index": 0,
//      "text": "The price for bubblegum decreased on thursday.",
//      "logprobs": null,
//      "finish_reason": "stop"
//    }
//  ]
//}

type CompletionChoice struct {
	Index        int    `json:"index"`
	Text         string `json:"text"`
	Logprobs     string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type CompletionResponse struct {
	Model   string             `json:"model"`
	Object  string             `json:"object"`
	ID      string             `json:"id"`
	Created int                `json:"created"`
	Choices []CompletionChoice `json:"choices"`
}

// {
//  "error": {
//    "code": "string",
//    "message": "string",
//    "param": "string",
//    "type": "string"
//  }
//}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

// {
//  "input": "This is a test.",
//  "user": "string",
//  "input_type": "query",
//  "model": "string",
//  "additionalProp1": {}
//}

type EmbeddingRequest struct {
	Input      string                 `json:"input"`
	User       string                 `json:"user"`
	InputType  string                 `json:"input_type"`
	Model      string                 `json:"model"`
	Additional map[string]interface{} `json:"additionalProp1"`
}

// {
//  "object": "string",
//  "model": "string",
//  "data": [
//    {
//      "index": 0,
//      "object": "string",
//      "embedding": [
//        0
//      ]
//    }
//  ],
//  "usage": {
//    "prompt_tokens": 0,
//    "total_tokens": 0
//  }
//}

type Data struct {
	Index     int     `json:"index"`
	Object    string  `json:"object"`
	Embedding []int64 `json:"embedding"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type EmbeddingResponse struct {
	Object string `json:"object"`
	Model  string `json:"model"`
	Data   []Data `json:"data"`
	Usage  Usage  `json:"usage"`
}

// {
//  "model": "gpt-35-turbo",
//  "messages": [
//    {
//      "role": "user",
//      "content": "Hello!"
//    }
//  ]
//}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// {
//  "id": "chatcmpl-123",
//  "object": "chat.completion",
//  "created": 1677652288,
//  "choices": [
//    {
//      "index": 0,
//      "message": {
//        "role": "assistant",
//        "content": "\n\nHello there, how may I assist you today?"
//      },
//      "finish_reason": "stop"
//    }
//  ],
//  "usage": {
//    "prompt_tokens": 9,
//    "completion_tokens": 12,
//    "total_tokens": 21
//  }
//}

type ChatChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ChatResponse struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []ChatChoice `json:"choices"`
	Usage   Usage        `json:"usage"`
}
