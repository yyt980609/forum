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

var Forum *gin.Engine

// SetUp 启动项目
func SetUp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Forum = gin.New()
	Forum.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true),
		ginzap.RecoveryWithZap(zap.L(), true),
		middleware.GinI18nLocalize(),
		middleware.JWTAuthMiddleware(),
	)
	Forum.GET("/", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		controller.ResponseSuccess(c, "Success")
	})
	// 注册
	Forum.POST("/register", controller.RegisterHandler)
	// 登陆
	Forum.POST("/login", controller.LoginHandler)
	// 404
	Forum.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})
	return Forum
}
