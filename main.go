package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AppName    string = "go-bookshelf-api"
	AppVersion string = "0.0.1"
)

func main() {
	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": AppName,
			"healthy": "true",
		})
	})

	router.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": AppName,
			"version": AppVersion,
		})
	})

	// listen on 0.0.0.0:8080 by default
	router.Run()
}
