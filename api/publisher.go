package api

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/rizkiamr/go-bookshelf-api/db/sqlc"
)

type createPublisherRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createPublisher(ctx *gin.Context) {
	var req createPublisherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal menambahkan Publisher. Mohon isi nama Publisher",
		})
		return
	}

	uuidObj := uuid.New()
	uuidStr := uuidObj.String()
	uuidWithoutDashes := strings.ReplaceAll(uuidStr, "-", "")

	arg := db.CreatePublisherParams{
		ID:   uuidWithoutDashes,
		Name: req.Name,
	}

	publisher, err := server.store.CreatePublisher(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Publisher berhasil ditambahkan",
		"data": map[string]string{
			"publisherId": publisher.ID,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": map[string][]db.Publisher{
			"publishers": publishers,
		},
	})
}

type getPublisherRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
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
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Publisher tidak ditemukan",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": map[string]db.Publisher{
			"publisher": publisher,
		},
	})
}

type updatePublisherRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) updatePublisher(ctx *gin.Context) {
	var req updatePublisherRequest

	id := ctx.Params.ByName("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui Publisher. Mohon isi nama publisher",
		})
		return
	}

	arg := db.UpdatePublisherParams{
		ID:   id,
		Name: req.Name,
	}

	_, err := server.store.UpdatePublisher(context.Background(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Gagal memperbarui publisher. Id tidak ditemukan",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Publisher berhasil diperbarui",
	})
}

type deletePublisherRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (server *Server) deletePublisher(ctx *gin.Context) {
	var req deletePublisherRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Publisher gagal dihapus. Id tidak valid",
		})
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

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Publisher berhasil dihapus",
	})
}
