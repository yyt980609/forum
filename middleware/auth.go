package middleware

import (
	"forum/common"
	"forum/controller"
	"forum/pkg/jwt"
	"forum/utils/forum_error"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取请求头中到token，Authorization:Bearer xxx.xx.xx
		authHeader := c.Request.Header.Get(common.TokenHeader)
		if authHeader == "" {
			controller.ResponseFailed(c, forum_error.CodeNoLogin)
			c.Abort()
			return
		}
		// 按空格分割，只取bearer后边到token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == common.Bearer) {
			controller.ResponseFailed(c, forum_error.CodeNoLogin)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		claim, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseFailed(c, forum_error.CodeNoLogin)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(common.UserId, claim.UserId)
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Next()
	}
}
