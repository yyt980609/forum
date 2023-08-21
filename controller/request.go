package controller

import (
	"forum/common"
	"forum/utils/forum_error"

	"github.com/gin-gonic/gin"
)

// GetCurrentUser 请求上下文中获取当前登陆用户
func GetCurrentUser(c *gin.Context) (int64, error) {
	userId, ok := c.Get(common.UserId)
	if !ok {
		return 0, forum_error.New(forum_error.CodeUserNotLogin)
	}
	return userId.(int64), nil
}
