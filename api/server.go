package api

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/rizkiamr/go-bookshelf-api/db/sqlc"
	ratelimit "github.com/rizkiamr/go-bookshelf-api/ratelimit"
	"github.com/rizkiamr/go-bookshelf-api/util"
)

// Server serves HTTP requests for our bookshelf service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func keyFunc(ctx *gin.Context) string {
	return ctx.ClientIP()
}

func errorHandler(ctx *gin.Context, info ratelimit.Info) {
	ctx.String(429, "Mau nge-DDoS deck? Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	config, err := util.LoadConfig("../.")
	if err != nil {
		panic(err)
	}

	// This makes it so each ip can only make 12 request per minute
	rateLimitQuota, err := strconv.Atoi(config.ServiceRateLimitQuotaPerMinute)
	if err != nil {
		panic(err)
	}

	rateLimitStore := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: uint(rateLimitQuota),
	})

	rateLimitMiddleware := ratelimit.RateLimiter(rateLimitStore, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	router.Use(rateLimitMiddleware)

	basePath := router.Group("/" + config.ServiceBasePath)
	basePath.GET("/healthz", server.healthzRoutes)
	basePath.GET("/version", server.versionRoutes)

	internal := basePath.Group("/internal")
	internal.GET("/metrics", server.addPrometheusHandler())

	v1 := basePath.Group("/v1")
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
