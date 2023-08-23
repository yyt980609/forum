package controller

import (
	"errors"
	fError "forum/utils/forum_error"
	"net/http"

	"github.com/gin-gonic/gin"
)

const SUCCESS = 1

const Failed = 0

// Response 响应
type Response struct {
	Status int         `json:"status"`
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

// ResponseFailed 失败响应
func ResponseFailed(c *gin.Context, code int) {
	r := &Response{
		Status: Failed,
		Code:   code,
		Msg:    fError.GetMsg(c, code),
		Data:   nil,
	}
	c.JSON(http.StatusOK, r)
}

// ResponseFailedWithMsg 自定义错误信息的失败响应
func ResponseFailedWithMsg(c *gin.Context, code int, msg interface{}) {
	r := &Response{
		Status: Failed,
		Code:   code,
		Msg:    msg,
		Data:   nil,
	}
	c.JSON(http.StatusOK, r)
}

// ResponseSuccess 成功响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	r := &Response{Status: SUCCESS, Code: 0000, Msg: nil, Data: data}
	c.JSON(http.StatusOK, r)
}

// BuildResponse 组装失败响应
func BuildResponse(c *gin.Context, data interface{}, err error) {
	var forumError *fError.ForumError
	if err == nil {
		ResponseSuccess(c, data)
	} else if errors.As(err, &forumError) {
		ResponseFailed(c, forumError.Code)
	} else {
		ResponseFailed(c, fError.CodeSystemError)
	}
}
