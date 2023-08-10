package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code int) {
	r := &Response{
		Code: code,
		Msg:  GetMsg(c),
		Data: nil,
	}
	c.JSON(http.StatusOK, r)
}
