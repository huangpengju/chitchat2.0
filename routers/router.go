package routers

import (
	"net/http"

	"chitchat2.0/pkg/setting"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func InitRouter() *gin.Engine {
	// 设置应用的模式【debug 开发 release 发布】
	gin.SetMode(setting.RunMode) // 自定义包 setting 中的 RunMode 存放配置文件中 RUN_MODE 的值，现在代替gin.SetMode(gin.DebugMode)

	// 创建不带中间件的路由
	r := gin.New()
	// Swagger API 文档的路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	user := UserInfo{
		Name:   "用户：张三",
		Gender: "男",
		Age:    18,
	}
	r.Static("/public", "./public")
	// 加载 templates 中所有模板文件
	// 使用不同目录下名称相同的模板，注意：一定要放在配置路由之前才行
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/chitchat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "前台-首页",
			"user":  user,
		})
	})
	user1 := UserInfo{
		Name:   "管理员：李四",
		Gender: "男",
		Age:    18,
	}
	r.GET("/admin", func(ctx *gin.Context) {
		// time.Sleep(5 * time.Second)
		// ctx.String(http.StatusOK, "你好a")
		ctx.HTML(http.StatusOK, "admin/index.html", map[string]interface{}{
			"title": "后台-首页",
			"user":  user1,
		})
	})
	return r
}
