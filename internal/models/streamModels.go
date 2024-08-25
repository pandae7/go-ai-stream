package models

import "github.com/tmc/langchaingo/llms"

type queryRequest struct {
}

type queryResponse struct {
}

// For a Chat Question and Chat response
// Only one Object is defined
type ChatModel struct {
	ChatType      string // denotes "Query" or "Response"
	LanguageCode  string
	LanguageModel string // denotes which LLM is being used, eg:- GPT, GEMINI, CLAUDE etc
	Content       llms.ContentResponse
}

// ChatContext
type ChatContext struct {
	id        string
	sessionId string
	chatList  []ChatModel
}
