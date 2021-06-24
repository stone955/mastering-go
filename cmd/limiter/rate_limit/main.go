package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

const defaultRate = 10

var rateLimit = ratelimit.New(defaultRate)

// 在 gin.HandlerFunc 加入限流逻辑
func leakyBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(c *gin.Context) {
		now := rateLimit.Take()
		log.Println(now.Sub(prev)) // 为了打印间隔时间
		prev = now
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", leakyBucket(), func(c *gin.Context) {
		c.JSON(http.StatusOK, true)
	})
	_ = r.Run(":8080")

	// 使用压测工具ab: ab -n 10 -c 2 http://127.0.0.1:8080/ping
}
