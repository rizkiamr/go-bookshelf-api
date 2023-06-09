package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkiamr/go-bookshelf-api/util"
)

func (server *Server) versionRoutes(ctx *gin.Context) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name":    config.ServiceName,
		"version": config.ServiceVersion,
	})
}
