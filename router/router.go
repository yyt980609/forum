package router

import (
	"forum/controller"
	"forum/utils/snowflake"
	"net/http"
	"strconv"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true), ginzap.RecoveryWithZap(zap.L(), true))
	r.GET("/", func(c *gin.Context) {
		id, _ := snowflake.GetID()
		c.String(http.StatusOK, "Ok"+strconv.Itoa(int(id)))
	})
	r.POST("/sign", controller.SignHandler)

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})
	return r
}
