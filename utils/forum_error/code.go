package forum_error

import (
	"strconv"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

const (
	// CodeSystemError 系统异常
	CodeSystemError = 1000 + iota
	CodeNoLogin
)
const (
	// CodeBusinessError 业务异常
	CodeBusinessError = 2000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeUserNotLogin
)

// GetMsg 解析错误信息
func GetMsg(c *gin.Context, code int) string {
	message, err := ginI18n.GetMessage(c, strconv.Itoa(code))
	if err != nil {
		//todo 获取上下文语言信息，返回内部异常
		return ""
	}
	return message
}
