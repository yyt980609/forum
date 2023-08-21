package controller

import (
	"errors"
	"forum/common"
	"forum/models"
	"forum/pkg/validate"
	"forum/service"
	"forum/utils/forum_error"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// RegisterHandler 注册处理函数
func RegisterHandler(c *gin.Context) {
	// 参数校验
	var p models.RegisterForm
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if ok {
			ResponseFailedWithMsg(c, forum_error.CodeInvalidPassword, validate.RemoveTopStruct(errs.Translate(validate.Trans)))
		} else {
			ResponseFailed(c, forum_error.CodeInvalidParam)
		}
		return
	}
	err := service.Register(&p)
	BuildFailedResponse(c, nil, err)
}

// LoginHandler 登陆处理函数
func LoginHandler(c *gin.Context) {
	p := new(models.LoginForm)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if ok {
			ResponseFailedWithMsg(c, forum_error.CodeInvalidPassword, validate.RemoveTopStruct(errs.Translate(validate.Trans)))
		} else {
			ResponseFailed(c, forum_error.CodeInvalidParam)
		}
		return
	}
	aToken, rToken, err := service.Login(p)
	token := make(map[string]string)
	token[common.AToken] = aToken
	token[common.RToken] = rToken
	BuildFailedResponse(c, token, err)
}
