package middleware

import (
	"github.com/892294101/jxlife/server/common"
	"github.com/892294101/jxlife/server/response"
	"github.com/gin-gonic/gin"
)

// JwtAuth JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Failed("未登录或非法访问", c)
			c.Abort()
			return
		}
		if err := common.VerifyToken(token); err != nil {
			response.Failed("登录已过期，请重新登录", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
