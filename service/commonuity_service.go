package service

import (
	"forum/dao"
	"forum/models"
	fError "forum/utils/forum_error"
	"strconv"

	"go.uber.org/zap"
)

// GetCommunityList 查询社区内容
func GetCommunityList() (data []*models.Community, err error) {
	data, err = dao.SelectCommunityList()
	if err != nil {
		zap.L().Error("SelectCommunityList failed.", zap.Error(err))
		return nil, fError.New(fError.CodeNoData)
	}
	return
}

// GetCommunityDetail 获取社区详情
func GetCommunityDetail(idStr string) (data *models.Community, err error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, fError.New(fError.CodeInvalidParam)
	}
	data, err = dao.SelectCommunityById(id)
	if err != nil {
		zap.L().Error("SelectCommunityById failed.", zap.Error(err))
		return nil, fError.New(fError.CodeNoData)
	}
	return
}
