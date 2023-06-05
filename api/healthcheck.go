package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkiamr/go-bookshelf-api/constant"
)

func (server *Server) healthzRoutes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":    constant.AppName,
		"healthy": "true",
	})
}
