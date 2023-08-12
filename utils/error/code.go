package error

import (
	"strconv"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

const (
	CodeSystemError = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
)

func GetMsg(c *gin.Context, code int) string {
	message, err := ginI18n.GetMessage(c, strconv.Itoa(code))
	if err != nil {
		//todo 获取上下文语言信息，返回内部异常
		return ""
	}
	return message
}
