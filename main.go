package main

import (
	"net/http"

	_ "chitchat2.0/docs"
	"chitchat2.0/routers"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title ChitChat API
// @version 2.0
// @description 这里是 ChitChat 论坛项目的 API 文档。
// @termsOfService http://swagger.io/terms/

// @contact.name 作者：黄鹏举
// @contact.url https://huangpengju.github.io/

// @license.name	Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// BasePath /api/v1

// @schemes http https
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				正在使用的安全定义的描述
func main() {
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := routers.InitRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	s.ListenAndServe()
}
