package dao

import (
	"forum/models"
	"forum/pkg/md5"
	"forum/pkg/mysql"

	"go.uber.org/zap"
)

// InsertUser 插入一条用户记录
func InsertUser(user *models.User) (err error) {
	db := mysql.GetDB()
	user.Password = md5.EncryptPassword(user.Password)
	tx := db.Table("user").Create(user)
	if tx.Error != nil {
		zap.L().Info("Insert user.")
		err = tx.Error
		return
	} else {
		return
	}
}

// SelectUser 查询用户信息
func SelectUser(p *models.User) (r *models.User, err error) {
	db := mysql.GetDB()
	result := db.Table("user").Where(p).First(&r)
	err = result.Error
	return
}
