package controller

import (
	"errors"
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
		ParamValidFailedResponse(c, err)
		return
	}
	err := service.Register(&p)
	BuildResponse(c, nil, err)
}

// LoginHandler 登陆处理函数
func LoginHandler(c *gin.Context) {
	p := new(models.LoginForm)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if ok {
			ResponseFailedWithMsg(c, forum_error.CodeInvalidParam, validate.RemoveTopStruct(errs.Translate(validate.Trans)))
		} else {
			ResponseFailed(c, forum_error.CodeSystemError)
		}
		return
	}
	user, err := service.Login(p)
	BuildResponse(c, user, err)
}
