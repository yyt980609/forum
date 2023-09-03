package dao

import (
	"forum/models"
	"forum/pkg/mysql"
)

func SelectCommunityList() (data []*models.Community, err error) {
	db := mysql.GetDB()
	result := db.Table("community").Find(&data)
	err = result.Error
	return
}

func SelectCommunityById(id int64) (data *models.Community, err error) {
	db := mysql.GetDB()
	result := db.Table("community").Where("community_id = ?", id).First(&data)
	err = result.Error
	return
}
