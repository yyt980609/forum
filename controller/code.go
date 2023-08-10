package controller

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

const (
	Success = 1000 + iota
	InvalidParam
	UserExist
	UserNotExist
	InvalidPassword
)

func GetMsg(c *gin.Context) string {
	message, err := ginI18n.GetMessage(c, "Success")
	if err != nil {
		return ""
	}
	return message
}
