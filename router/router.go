package router

import (
	"forum/controller"
	"forum/middleware"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true),
		ginzap.RecoveryWithZap(zap.L(), true),
		middleware.GinI18nLocalize(),
	)
	r.GET("/1", func(c *gin.Context) {
		controller.ResponseError(c, 1)
	})
	// 注册
	r.POST("/register", controller.RegisterHandler)
	// 登陆
	r.POST("/login", controller.LoginHandler)
	// 404
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})
	return r
}
