package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkiamr/go-bookshelf-api/constant"
)

func (server *Server) versionRoutes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":    constant.AppName,
		"version": constant.AppVersion,
	})
}
