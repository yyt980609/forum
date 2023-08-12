package controller

import (
	"forum/utils/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

const SUCCESS = 1

const Failed = 0

type Response struct {
	Status int         `json:"status"`
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func ResponseFailed(c *gin.Context, code int) {
	r := &Response{
		Status: Failed,
		Code:   code,
		Msg:    error.GetMsg(c, code),
		Data:   nil,
	}
	c.JSON(http.StatusOK, r)
}

func ResponseFailedWithMsg(c *gin.Context, code int, msg interface{}) {
	r := &Response{
		Status: Failed,
		Code:   code,
		Msg:    msg,
		Data:   nil,
	}
	c.JSON(http.StatusOK, r)
}
