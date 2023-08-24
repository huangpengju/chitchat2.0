package routers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"chitchat2.0/pkg/setting"
	"chitchat2.0/routers/api"
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
	// 全局中间件 Logger 中间件将日志写入 gin.DefaultWriter
	r.Use(gin.Logger())
	// Recovery 中间件会 revover 恢复 任何 panic
	r.Use(gin.Recovery())

	// /public 表示路由，setting.Staic 表示路径 ./public
	r.Static("/static", setting.Static)

	// Swagger API 文档的路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.SetFuncMap(template.FuncMap{
		"formatDate": formatAsDate,
	})
	// 加载 templates 中所有模板文件
	// 使用不同目录下名称相同的模板，注意：一定要放在配置路由之前才行
	r.LoadHTMLGlob("templates/**/*")
	// 设置路由组 v1 ，处理登录
	// 用户登录界面
	r.GET("/login", api.Login)
	r.POST("/authenticate", api.Authenticate)
	r.GET("/signup", api.Signup)
	r.POST("/signup_account", api.SignupAccount)
	// v1 := r.Group("api/v1")
	// {

	// }
	// 登录（验证用户名和密码），并签发 Token
	// r.GET("/login", api.GetAuth)

	users := []UserInfo{
		{Uuid: 1, Name: "李四1", Gender: "男", Age: 17, CreatedAtDate: "2023-08-22", NumReplies: 1, Topic: "前台1号---张三"},
		{Uuid: 2, Name: "王二2", Gender: "男", Age: 18, CreatedAtDate: "2023-08-23", NumReplies: 2, Topic: "前台2号---王二"},
		{Uuid: 3, Name: "麻子3", Gender: "男", Age: 19, CreatedAtDate: "2023-08-24", NumReplies: 3, Topic: "前台3号---麻子"}}

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
