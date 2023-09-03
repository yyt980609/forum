package controller

import (
	"forum/models"
	"forum/service"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	var p models.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Create post with invalid param", zap.Error(err))
		ParamValidFailedResponse(c, err)
		return
	}
	user, err := GetCurrentUser(c)
	if err != nil {
		BuildResponse(c, nil, err)
	}
	p.AuthorId = user
	err, id := service.CreatePost(&p)
	BuildResponse(c, id, err)
}

// GetPostDetailHandler 查看帖子详情
func GetPostDetailHandler(c *gin.Context) {
	id := c.Param("id")
	data, err := service.GetPostDetail(id)
	BuildResponse(c, data, err)
}

// GetPostListHandler 分页查询帖子
func GetPostListHandler(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Param("currentPage"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	orderParam := string(c.Param("order"))
	var order string
	if strings.Contains(orderParam, "score") {
		order = "SCORE DESC"
	} else {
		order = "create_time DESC"
	}
	data, err := service.GetPostList(currentPage, pageSize, order)
	BuildResponse(c, data, err)
}
