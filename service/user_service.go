package service

import (
	"forum/dao"
	"forum/models"
	"forum/pkg/jwt"
	"forum/pkg/md5"
	"forum/pkg/snowflake"
	fError "forum/utils/forum_error"

	"go.uber.org/zap"
)

// Register 注册逻辑
func Register(p *models.RegisterForm) (err error) {
	u, err := dao.SelectUser(&models.User{Username: p.Username})
	if err == nil {
		return fError.New(fError.CodeUserExist)
	}
	id, e := snowflake.GetID()
	if e != nil {
		zap.L().Error(e.Error())
		return e
	}
	u = &models.User{UserID: id, Username: p.Username, Password: p.Password}
	return dao.InsertUser(u)
}

// Login 登陆逻辑
func Login(p *models.LoginForm) (aToken, rToken string, err error) {
	user, err := dao.SelectUser(&models.User{Username: p.Username, Password: md5.EncryptPassword(p.Password)})
	if err != nil {
		return "", "", fError.New(fError.CodeUserNotExist)
	}
	// 生成JwtToken
	return jwt.GenToken(user.UserID, user.Username)
}
