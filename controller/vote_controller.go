package controller

import (
	"forum/models"
	"forum/service"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// PostVoteHandler 投票
func PostVoteHandler(c *gin.Context) {
	v := new(models.VoteData)
	if err := c.ShouldBindJSON(&v); err != nil {
		zap.L().Error("Vote with invalid param", zap.Error(err))
		ParamValidFailedResponse(c, err)
		return
	}
	userId, err := GetCurrentUser(c)
	if err != nil {
		BuildResponse(c, nil, err)
	}
	err = service.Vote(userId, v)
	BuildResponse(c, nil, err)
}
