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

type createAuthorRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createAuthor(ctx *gin.Context) {
	var req createAuthorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal menambahkan Author. Mohon isi nama Author",
		})
		return
	}

	uuidObj := uuid.New()
	uuidStr := uuidObj.String()
	uuidWithoutDashes := strings.ReplaceAll(uuidStr, "-", "")

	arg := db.CreateAuthorParams{
		ID:   uuidWithoutDashes,
		Name: req.Name,
	}
	author, err := server.store.CreateAuthor(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Author berhasil ditambahkan",
		"data": map[string]string{
			"authorId": author.ID,
		},
	})
}

type listAuthorRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAuthors(ctx *gin.Context) {
	var req listAuthorRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAuthorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	authors, err := server.store.ListAuthors(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": map[string][]db.Author{
			"authors": authors,
		},
	})
}

type getAuthorRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAuthor(ctx *gin.Context) {
	var req getAuthorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	author, err := server.store.GetAuthor(context.Background(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Author tidak ditemukan",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": map[string]db.Author{
			"author": author,
		},
	})
}

type updateAuthorRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) updateAuthor(ctx *gin.Context) {
	var req updateAuthorRequest

	id := ctx.Params.ByName("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui author. Mohon isi nama author",
		})
		return
	}

	arg := db.UpdateAuthorParams{
		ID:   id,
		Name: req.Name,
	}

	_, err := server.store.UpdateAuthor(context.Background(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Gagal memperbarui author. Id tidak ditemukan",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Author berhasil diperbarui",
	})
}

type deleteAuthorRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteAuthor(ctx *gin.Context) {
	var req deleteAuthorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Author gagal dihapus. Id tidak valid",
		})
		return
	}

	err := server.store.DeleteAuthor(context.Background(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Author gagal dihapus. Id tidak ditemukan",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Author berhasil dihapus",
	})
}
