package api

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/rizkiamr/go-bookshelf-api/db/sqlc"
)

type createBookRequest struct {
	Name        string `json:"name" binding:"required"`
	Year        int32  `json:"year"`
	AuthorID    int64  `json:"author_id"`
	Summary     string `json:"summary"`
	PublisherID int64  `json:"publisher_id"`
	PageCount   int32  `json:"pageCount"`
	ReadPage    int32  `json:"readPage"`
	Reading     bool   `json:"reading"`
}

func (server *Server) createBook(ctx *gin.Context) {
	var req createBookRequest

	readFinished := false

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal menambahkan buku. Mohon isi nama buku",
		})
		return
	}

	if req.ReadPage > req.PageCount {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
		})
		return
	}

	if req.PageCount == req.ReadPage {
		readFinished = true
	}

	tz, _ := time.LoadLocation("Etc/UTC")

	arg := db.CreateBookParams{
		Name:        req.Name,
		Year:        req.Year,
		AuthorID:    req.AuthorID,
		Summary:     req.Summary,
		PublisherID: req.PublisherID,
		PageCount:   req.PageCount,
		ReadPage:    req.ReadPage,
		Finished:    readFinished,
		Reading:     req.Reading,
		UpdatedAt:   sql.NullTime{Time: time.Now().In(tz), Valid: true},
	}

	book, err := server.store.CreateBook(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Buku berhasil ditambahkan",
		"data": map[string]int64{
			"bookId": book.ID,
		},
	})
}

type listBookRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listBooks(ctx *gin.Context) {
	var req listBookRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListBooksParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	books, err := server.store.ListBooks(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": map[string][]db.Book{
			"books": books,
		},
	})
}

type getBookRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBook(ctx *gin.Context) {
	var req getBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.store.GetBook(context.Background(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Buku tidak ditemukan",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"status": "success",
		"data": map[string]db.Book{
			"book": book,
		},
	})
}

type updateBookRequest struct {
	Name        string `json:"name" binding:"required"`
	Year        int32  `json:"year"`
	AuthorID    int64  `json:"author_id"`
	Summary     string `json:"summary"`
	PublisherID int64  `json:"publisher_id"`
	PageCount   int32  `json:"pageCount"`
	ReadPage    int32  `json:"readPage"`
	Reading     bool   `json:"reading"`
}

func (server *Server) updateBook(ctx *gin.Context) {
	var req updateBookRequest

	id, _ := strconv.ParseInt(ctx.Params.ByName("id"), 0, 64)

	readFinished := false

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui buku. Mohon isi nama buku",
		})
		return
	}

	if req.ReadPage > req.PageCount {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
		})
		return
	}

	if req.PageCount == req.ReadPage {
		readFinished = true
	}

	tz, _ := time.LoadLocation("Etc/UTC")

	arg := db.UpdateBookParams{
		ID:          id,
		Name:        req.Name,
		Year:        req.Year,
		AuthorID:    req.AuthorID,
		Summary:     req.Summary,
		PublisherID: req.PublisherID,
		PageCount:   req.PageCount,
		ReadPage:    req.ReadPage,
		Finished:    readFinished,
		Reading:     req.Reading,
		UpdatedAt:   sql.NullTime{Time: time.Now().In(tz), Valid: true},
	}

	_, err := server.store.UpdateBook(context.Background(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Gagal memperbarui buku. Id tidak ditemukan",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Buku berhasil diperbarui",
	})
}

type deleteBookRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteBook(ctx *gin.Context) {
	var req deleteBookRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Buku gagal dihapus. Id tidak valid",
		})
		return
	}
	err := server.store.DeleteBook(context.Background(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Buku gagal dihapus. Id tidak ditemukan",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Buku berhasil dihapus",
	})
}
