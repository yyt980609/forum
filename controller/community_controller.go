package controller

import (
	"forum/service"

	"github.com/gin-gonic/gin"
)

// CommunityHandler 获取所以社区
func CommunityHandler(c *gin.Context) {
	data, err := service.GetCommunityList()
	BuildResponse(c, data, err)
}

// CommunityDetailHandler 获取社区详情
func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	data, err := service.GetCommunityDetail(idStr)
	BuildResponse(c, data, err)
}
