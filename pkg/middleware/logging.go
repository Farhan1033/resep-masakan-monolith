package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	file, err := os.OpenFile("logs/api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open API log file: %v", err)
	}

	logger := log.New(file, "[API] ", log.Ldate|log.Ltime)

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Printf("%s %s | %d | %v | %s",
			method, path, status, latency, ip)
	}
}
