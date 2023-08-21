package routers

import (
	"net/http"

	"chitchat2.0/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 设置应用的模式【debug 开发 release 发布】
	gin.SetMode(setting.RunMode) // 自定义包 setting 中的 RunMode 存放配置文件中 RUN_MODE 的值，现在代替gin.SetMode(gin.DebugMode)

	// 创建不带中间件的路由
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好")
	})
	return r
}
