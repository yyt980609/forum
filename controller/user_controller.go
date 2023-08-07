package controller

import (
	"forum/service"

	"github.com/gin-gonic/gin"
)

func SignHandler(c *gin.Context) {
	// 参数校验
	service.Sign()
}
