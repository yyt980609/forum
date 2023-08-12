package controller

import (
	"errors"
	"forum/models"
	"forum/service"
	"forum/utils/error"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	// 参数校验
	var p models.RegisterForm
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if ok {
			ResponseFailedWithMsg(c, error.CodeInvalidPassword, removeTopStruct(errs.Translate(trans)))
		} else {
			ResponseFailed(c, error.CodeInvalidParam)
		}
		return
	}
	err := service.Register(&p)
	var forumError *error.ForumError
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	} else if errors.As(err, &forumError) {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "err",
		})
	}

}

func LoginHandler(c *gin.Context) {
	p := new(models.LoginForm)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if ok {
			ResponseFailedWithMsg(c, error.CodeInvalidPassword, removeTopStruct(errs.Translate(trans)))
		} else {
			ResponseFailed(c, error.CodeInvalidParam)
		}
		return
	}
	err := service.Login(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
