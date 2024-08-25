package langchain

// OpenAI LLM Provider, Handles all communications with OpenAILLM

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"log"

	"github.com/tmc/langchaingo/llms/openai"
)

// List of OpenAI Functions integrated with Langchain -
// 1. Call -> Generates using Prompt string
// 2. CreateEmbedding
// 3. GenerateContent -> Takes in an array of []llms.MessageContent

// List to Contain all the call Options for running
type LLMCallOptions struct {
	options llms.CallOptions
}

type OpenAIConstructor struct {
	model string
}

// Should Create a Model for generic LLM Response

func (oc *OpenAIConstructor) runPrompt(ctx context.Context, prompt string) openai.ResponseFormat {
	openaiClient, err := openai.New(openai.WithModel(oc.model))
	if err != nil {
		log.Panic("OpenAI Client Provider failed")
	}
	response, err := openaiClient.Call(ctx, prompt)
	return openai.ResponseFormat{Type: response}
}

func (oc *OpenAIConstructor) generateContent(ctx context.Context, prompt string) *llms.ContentResponse {
	openaiClient, err := openai.New(openai.WithModel(oc.model))
	if err != nil {
		log.Panic("OpenAI Client Provider failed")
	}
	initialMessage := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	}

	completion, err := openaiClient.GenerateContent(ctx, initialMessage, llms.WithStreamingFunc(
		func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	return completion
}
