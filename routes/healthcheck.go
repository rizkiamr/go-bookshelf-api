package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkiamr/go-bookshelf-api/constant"
)

func addHealthzRoutes() {
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    constant.AppName,
			"healthy": "true",
		})
	})
}
