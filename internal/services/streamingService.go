package services

import (
	//"textStreaming/internal/models"
	"github.com/gin-gonic/gin"
	"textStreaming/internal/managers"
)

type StreamConstructor struct {
	StreamManager managers.StreamManager
	Prompt        string
}

func (sc *StreamConstructor) Chat(c *gin.Context) {
	chatResponse := sc.StreamManager.HandleChat(sc.Prompt)
	c.JSON(200, chatResponse)
}
