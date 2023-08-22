package main

import (
	"fmt"
	"log"
	"syscall"

	_ "chitchat2.0/docs"
	"chitchat2.0/pkg/setting"
	"chitchat2.0/routers"
	"github.com/fvbock/endless"
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
	// 普通启动
	// router := routers.InitRouter()
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
	// 优雅地重启
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("现在的 pid 是 %d", syscall.Getpid())

	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("服务器错误：%v", err)
	}
}

/* 优雅地关机
{
	go func() {
		// 服务连接
		if err := s.ListenAndServe(); err != nil {
			// logging.Warn("Listen:%s\n", err)
			log.Printf("Listen:%s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统的SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify 把收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// logging.Info("关闭服务器...")
	log.Println("关闭服务器...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("服务器关闭：", err)
	}
	log.Println("服务器正在退出...")
}
*/
