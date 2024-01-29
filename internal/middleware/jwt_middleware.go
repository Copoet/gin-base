package middleware

import (
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		//获取token header
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			code = util.PublicAuthError
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  util.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		//截取token
		tokenString := strings.TrimPrefix(token, "Bearer ")
		// 校验token
		_, err := util.VerifyToken(tokenString)
		if err != nil {
			code = util.PublicAuthError
			data = nil
			c.JSON(http.StatusUnauthorized, gin.H{
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
