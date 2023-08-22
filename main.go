package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "chitchat2.0/docs"
	"chitchat2.0/pkg/setting"
	"chitchat2.0/routers"
)

// 设置通用 API 注释

// @title ChitChat API
// @version 2.0
// @description Go Web 实践项目 —— ChitChat2.0 API 文档。
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
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			// logging.Warn("Listen:%s\n", err)
			log.Printf("Listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// logging.Info("关闭服务器...")
	log.Println("关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("服务器关闭：", err)
	}
	log.Println("服务器正在退出...")
}
