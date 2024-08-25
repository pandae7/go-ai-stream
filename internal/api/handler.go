package api

import (
	"fmt"
	"textStreaming/internal/managers"

	"github.com/gin-gonic/gin"
	"textStreaming/internal/services"
)

// Should handle following things -
// Session token and etc
// Internal methods to talk to inference providers
// Session Stream API
// post question API
// fetch Context and also application to store it

func allHandlers() {

	router := gin.Default()
	streamManager := managers.StreamManager{}
	streamClient := services.StreamConstructor{
		StreamManager: streamManager,
		Prompt:        "Hello, You are a good Assistant",
	}
	router.GET("/", homePage)
	router.GET("/chat", streamClient.Chat)
}

func homePage(c *gin.Context) {
	fmt.Println("Hello World")
}
