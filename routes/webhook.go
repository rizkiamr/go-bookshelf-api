package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	helloWorldWebhookProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "webhook_hello_world_hit_total",
		Help: "The total number of hello-world webhook got invoked.",
	})

	helloWorldWebhookCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_hello_world_webhook_invocation_count", // metric name
		Help: "Number of hello-world webhook invocation.",
	},
		[]string{"status"})

	helloWorldWebhookLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_hello_world_webhook_invocation_latency_seconds",
		Help:    "Latency of hello-world webhook invocation in second.",
		Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
	},
		[]string{"status"})
)

func addWebhookRoutes(rg *gin.RouterGroup) {
	webhooks := rg.Group("/webhooks")

	webhooks.GET("/hello-world", helloWorldWebhookFunc)
}

func helloWorldWebhookFunc(c *gin.Context) {
	var status string
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		helloWorldWebhookLatency.WithLabelValues(status).Observe(v)
	}))

	// do something here
	_, err := fmt.Println("Hello, World!")

	if err != nil {
		status = "error"

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "something went wrong on our side",
		})

		helloWorldWebhookCounter.WithLabelValues(status).Inc()
		timer.ObserveDuration()
	}

	status = "success"
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Hello, World!",
	})

	helloWorldWebhookProcessed.Inc()
	helloWorldWebhookCounter.WithLabelValues(status).Inc()
	timer.ObserveDuration()
}
