package langchain

// Manager to handle integration with Langchain library
// Langchain is only used for it's abstractions between almost all the LLMs there are
// The Chaining of Langchain is usually for threading the output from Chat Models to Image Generation Models or something like that

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
)

type LangchainConstructor struct {
	LLM string
}

// This function ideally should only return the Client based on the LLM decided
// Generating Content or running with a single Prompt is a next step? Right?
func (lc *LangchainConstructor) LangchainClient(prompt string) *llms.ContentResponse {
	fmt.Println("initialization of langchainClient")
	fmt.Println(lc.LLM)
	fmt.Println(prompt)

	if lc.LLM == "openAI" {
		ctx := context.Background()
		openAIClient := OpenAIConstructor{
			model: "gpt-3.5-turbo-0125",
		}
		return openAIClient.generateContent(ctx, prompt)
	}
	return nil

}
