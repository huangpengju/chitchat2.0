package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// 创建不带中间件的路由
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好")
	})
	return r
}
