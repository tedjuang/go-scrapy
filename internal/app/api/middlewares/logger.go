package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware that logs request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		t := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(t)

		// Log request details
		log.Printf("[API] %s | %d | %s | %s | %v",
			c.Request.Method,
			c.Writer.Status(),
			c.Request.URL.Path,
			c.ClientIP(),
			latency,
		)
	}
}
