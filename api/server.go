package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/rizkiamr/go-bookshelf-api/db/sqlc"
)

// Server serves HTTP requests for our bookshelf service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/healthz", server.healthzRoutes)
	router.GET("/version", server.versionRoutes)

	internal := router.Group("/internal")
	internal.GET("/metrics", server.addPrometheusHandler())

	v1 := router.Group("/v1")
	v1.POST("/authors", server.createAuthor)
	v1.GET("/authors/:id", server.getAuthor)
	v1.GET("/authors", server.listAuthors)
	v1.DELETE("/authors/:id", server.deleteAuthor)
	v1.PUT("/authors/:id", server.updateAuthor)

	v1.POST("/publishers", server.createPublisher)
	v1.GET("/publishers/:id", server.getPublisher)
	v1.GET("/publishers", server.listPublishers)
	v1.DELETE("/publishers/:id", server.deletePublisher)
	v1.PUT("/publishers/:id", server.updatePublisher)

	v1.POST("/books", server.createBook)
	v1.GET("/books", server.listBooks)
	v1.GET("/books/:id", server.getBook)
	v1.PUT("/books/:id", server.updateBook)
	v1.DELETE("/books/:id", server.deleteBook)

	addWebhookRoutes(v1)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func deleteOkResponse(id int64) gin.H {
	return gin.H{
		"id":      id,
		"status":  "ok",
		"message": "successfully deleted",
	}
}
