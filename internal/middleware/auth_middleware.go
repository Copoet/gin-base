package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// LoggerMiddleware 是一个记录请求处理时间的中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求处理前的时间
		t := time.Now()

		// 调用c.Next()后，请求的处理就会进入到下一个阶段
		// 这可以是具体的请求处理函数
		c.Next()

		// 请求处理后的时间
		latency := time.Since(t)
		log.Printf("请求 %s 耗时 %v", c.Request.URL.Path, latency)
	}
}
