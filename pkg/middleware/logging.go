package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	logDir := "/tmp/logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Println("Cannot create /tmp/logs folder, fallback to stdout")
	}

	filePath := logDir + "/api.log"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Cannot open /tmp/logs/api.log, fallback to stdout")
		file = os.Stdout
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

		logger.Printf("%s %s | %d | %v | %s", method, path, status, latency, ip)
	}
}
