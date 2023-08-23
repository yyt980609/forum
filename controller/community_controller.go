package controller

import (
	"forum/service"

	"github.com/gin-gonic/gin"
)

func CommunityHandler(c *gin.Context) {
	data, err := service.GetCommunityList()
	BuildResponse(c, data, err)
}
