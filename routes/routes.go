package routes

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true), ginzap.RecoveryWithZap(zap.L(), true))
	r.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "ok")
	})
	return r
}
