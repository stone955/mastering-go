package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const (
	defaultInterval = 100 * time.Millisecond
	defaultTokens   = 10
)

var rateLimit = DefaultLimiter()

// DefaultLimiter ...
func DefaultLimiter() *rate.Limiter {
	return rate.NewLimiter(rate.Every(defaultInterval), defaultTokens)
}

// 在 gin.HandlerFunc 加入限流逻辑
func tokenBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		if rateLimit.Allow() {
			c.String(http.StatusOK, "rate limit,Drop")
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", tokenBucket(), func(c *gin.Context) {
		c.JSON(http.StatusOK, true)
	})
	_ = r.Run(":8080")

	// 使用压测工具ab: ab -n 10 -c 2 http://127.0.0.1:8080/ping
}
