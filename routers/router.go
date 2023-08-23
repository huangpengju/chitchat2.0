package routers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"chitchat2.0/pkg/setting"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type UserInfo struct {
	Name          string
	Gender        string
	Age           int
	CreatedAtDate string
	NumReplies    int
	Topic         string
	Uuid          int
}

func InitRouter() *gin.Engine {
	// 设置应用的模式【debug 开发 release 发布】
	gin.SetMode(setting.RunMode) // 自定义包 setting 中的 RunMode 存放配置文件中 RUN_MODE 的值，现在代替gin.SetMode(gin.DebugMode)

	// 创建不带中间件的路由
	r := gin.New()
	// Swagger API 文档的路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	users := []UserInfo{
		{Uuid: 1, Name: "李四", Gender: "男", Age: 17, CreatedAtDate: "2023-08-22", NumReplies: 1, Topic: "前台---张三"},
		{Uuid: 2, Name: "王二", Gender: "男", Age: 18, CreatedAtDate: "2023-08-23", NumReplies: 2, Topic: "前台---王二"},
		{Uuid: 3, Name: "麻子", Gender: "男", Age: 19, CreatedAtDate: "2023-08-24", NumReplies: 3, Topic: "前台---麻子"}}
	r.SetFuncMap(template.FuncMap{
		"formatDate": formatAsDate,
	})
	r.Static("/public", setting.Static)

	// 加载 templates 中所有模板文件
	// 使用不同目录下名称相同的模板，注意：一定要放在配置路由之前才行
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/chitchat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/layout.html", gin.H{
			"state": 1,
			"users": users,
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
			"now":   time.Now(),
		})
	})
	return r
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d%02d", year, month, day)
}
