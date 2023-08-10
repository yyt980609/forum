package service

import (
	"errors"
	"forum/dao"
	"forum/models"
	"forum/utils/md5"
	"forum/utils/snowflake"

	"go.uber.org/zap"
)

func Register(p *models.RegisterForm) (err error) {
	u, err := dao.SelectUser(&models.User{Username: p.UserName})
	if err != nil {
		return errors.New("用户名已存在")
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
		return errors.New("用户不存在或密码错误")
	}
	return err
}
