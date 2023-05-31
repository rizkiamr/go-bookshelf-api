package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkiamr/go-bookshelf-api/constant"
)

func addHealthCheckRoutes(rg *gin.RouterGroup) {
	healthz := rg.Group("/healthz")

	healthz.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    constant.AppName,
			"healthy": "true",
		})
	})
}
