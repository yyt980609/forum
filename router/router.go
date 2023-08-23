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

// SetUp 启动项目
func SetUp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	forum := gin.New()
	forum.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true),
		ginzap.RecoveryWithZap(zap.L(), true),
		middleware.GinI18nLocalize(),
	)
	v1 := forum.Group("/api/v1")
	// v1.Use(middleware.JWTAuthMiddleware())
	// 主页
	v1.GET("/", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		controller.ResponseSuccess(c, "Success")
	})
	// 注册
	v1.POST("/register", controller.RegisterHandler)
	// 登陆
	v1.POST("/login", controller.LoginHandler)
	{
		v1.GET("/community", controller.CommunityHandler)

	}

	// 404
	forum.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})
	return forum
}
