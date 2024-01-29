package middleware

import (
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		//获取token header
		token := c.GetHeader("Authorization")
		//校验token
		if !util.ValidTokenExpress(token) {
			code = util.PublicAuthError
			data = nil
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  util.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()

	}
}
