package service

import (
	"forum/dao"
	"forum/models"
	fError "forum/utils/forum_error"

	"go.uber.org/zap"
)

// GetCommunityList 查询社区内容
func GetCommunityList() (data []*models.Community, err error) {
	data, err = dao.SelectCommunityList()
	if err != nil {
		zap.L().Error("GetCommunityList failed.", zap.Error(err))
		return nil, fError.New(fError.CodeNoData)
	}
	return
}
