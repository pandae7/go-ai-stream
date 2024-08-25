package managers

import (
	"textStreaming/internal/models"
	"textStreaming/pkg/langchain"
)

// Stream Manager will handle all the background processes of the text Stream.
// This Manager is responsible for handling issues with the stream like -
// 1. LLM Model errors
// 2. Network Issues
// 3. Handles switch between multiple models based on qualifying criteria
// 4. Same functions are exposed into two places - Service APIs and Client Calls

type StreamManager struct {
	//
}

func (sm *StreamManager) HandleChat(prompt string) models.ChatModel {
	langchainConstructor := langchain.LangchainConstructor{LLM: "openAI"}
	chatResponse := langchainConstructor.LangchainClient(prompt)
	return models.ChatModel{
		ChatType:      "Response",
		LanguageCode:  "en",
		LanguageModel: "OpenAI",
		Content:       *chatResponse,
	}
}
