package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkiamr/go-bookshelf-api/constant"
)

func addVersionRoutes(rg *gin.RouterGroup) {
	rg.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    constant.AppName,
			"version": constant.AppVersion,
		})
	})
}
