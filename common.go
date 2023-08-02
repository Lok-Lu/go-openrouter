package openrouter

const (
	GooglePalm2CodeChatBison = "google/palm-2-codechat-bison"
	GooglePalm2ChatBison     = "google/palm-2-chat-bison"
	OpenaiGpt35Turbo         = "openai/gpt-3.5-turbo"
	OpenaiGpt35Turbo16k      = "openai/gpt-3.5-turbo-16k"
	OpenaiGpt4               = "openai/gpt-4"
	OpenaiGpt432K            = "openai/gpt-4-32k"
	AnthropicClaude2         = "anthropic/claude-2"
	AnthropicClaudeInstantV1 = "anthropic/claude-instant-v1"
	MetaLlamaLlama213bChat   = "meta-llama/llama-2-13b-chat"
	MetaLlamaLlama270bChat   = "meta-llama/llama-2-70b-chat"
)

var enableModelsForEndpoints = map[string]bool{
	GooglePalm2CodeChatBison: true,
	GooglePalm2ChatBison:     true,
	OpenaiGpt35Turbo:         true,
	OpenaiGpt35Turbo16k:      true,
	OpenaiGpt4:               true,
	OpenaiGpt432K:            true,
	AnthropicClaude2:         true,
	AnthropicClaudeInstantV1: true,
	MetaLlamaLlama213bChat:   true,
	MetaLlamaLlama270bChat:   true,
}

func checkSupportsModel(model string) bool {
	return enableModelsForEndpoints[model]
}

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionRequest represents a request structure for chat completion API.
type ChatCompletionRequest struct {
	Model       string                  `json:"model"`
	Messages    []ChatCompletionMessage `json:"messages"`
	Stream      bool                    `json:"stream,omitempty"`
	Temperature *float32                `json:"temperature,omitempty"`
	TopP        *float32                `json:"top_p,omitempty"`
	TopK        *uint                   `json:"top_k,omitempty"`
}

type ChatCompletionChoice struct {
	Message      string `json:"message,omitempty"`
	FinishReason string `json:"finish_reason,omitempty"`
	Delta        any    `json:"delta,omitempty"`
	Content      string `json:"content,omitempty"`
	User         string `json:"user,omitempty"`
	Index        uint   `json:"index,omitempty"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
	ID      string                 `json:"id,omitempty"`
	Object  string                 `json:"object,omitempty"`
	Created int64                  `json:"created,omitempty"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   Usage                  `json:"usage,omitempty"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
