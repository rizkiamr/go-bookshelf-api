package api

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/rizkiamr/go-bookshelf-api/db/sqlc"
)

type createPublisherRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createPublisher(ctx *gin.Context) {
	var req createPublisherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	name := req.Name

	publisher, err := server.store.CreatePublisher(context.Background(), name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, publisher)
}

type getPublisherRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPublisher(ctx *gin.Context) {
	var req getPublisherRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	publisher, err := server.store.GetPublisher(context.Background(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, publisher)
}

type listPublisherRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPublishers(ctx *gin.Context) {
	var req listPublisherRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPublishersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	publishers, err := server.store.ListPublishers(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, publishers)
}

type deletePublisherRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deletePublisher(ctx *gin.Context) {
	var req deletePublisherRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeletePublisher(context.Background(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, deleteOkResponse(req.ID))
}

type updatePublisherRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) updatePublisher(ctx *gin.Context) {
	var req updatePublisherRequest

	id, _ := strconv.ParseInt(ctx.Params.ByName("id"), 0, 64)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePublisherParams{
		ID:   id,
		Name: req.Name,
	}

	publisher, err := server.store.UpdatePublisher(context.Background(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, publisher)
}
