package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addWebhookRoutes(rg *gin.RouterGroup) {
	webhooks := rg.Group("/webhooks")

	webhooks.GET("/hello-world", helloWorldWebhookFunc)
}

func helloWorldWebhookFunc(c *gin.Context) {
	// do something here
	_, err := fmt.Println("Hello, World!")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "something went wrong on our side",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
