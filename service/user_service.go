package service

import (
	"forum/dao"
	"forum/models"
	fError "forum/utils/error"
	"forum/utils/md5"
	"forum/utils/snowflake"

	"go.uber.org/zap"
)

func Register(p *models.RegisterForm) (err error) {
	u, err := dao.SelectUser(&models.User{Username: p.UserName})
	if err == nil {
		return fError.New(fError.CodeUserExist)
	}
	id, e := snowflake.GetID()
	if e != nil {
		zap.L().Error(e.Error())
		return e
	}
	u = &models.User{UserID: id, Username: p.UserName, Password: p.Password}
	return dao.InsertUser(u)
}

func Login(p *models.LoginForm) (err error) {
	_, err = dao.SelectUser(&models.User{Username: p.UserName, Password: md5.EncryptPassword(p.Password)})
	if err != nil {
		return fError.New(fError.CodeUserNotExist)
	}
	return err
}
