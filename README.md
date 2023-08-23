# ChitChat2.0 论坛
ChitChat2.0 仓库是中 ChitChat 项目的升级版 

## 项目实现
后端使用 Go + Gin + Gorm + JWT + Mysql  
前端使用 HTML5 + Vue
## 项目思路
1. 搭建项目
* Go + Gin 开始搭建项目
* 为已搭建的项目加上 Swagger，Swag 将 Go 的注释转换为 Swagger2.0 文档
* Swag 支持 Gin 框架，实现 Swag 与 Gin 集成

2. 配置文件
* 创建 ini 文件
* 读取 ini 中的内容信息，应用到 gin.SetMode 和记录日志

3. 记录日志
* 日志目录和日志文件：存在则打开日志文件，不存在则创建
* 日志的记录：初始化时先打开日志文件，调用日志记录函数时设置日志前缀，输出日志

4. 优雅地关机
* http.Server - Shutdown()
* 优雅关机就是服务端关机命令发出后不是立即关机，而是等待当前还在处理的请求全部处理完毕后再退出程序，是一种对客户端友好的关机方式。而执行Ctrl+C关闭服务端时，会强制结束进程导致正在访问的请求出现问题。

5. 优雅地重启
* 使用 fvbock/endless 来替换默认的 ListenAndServe 启动服务来实现。
* 由于开发环境是 Windows ，且参考资料中需要使用命令： kill -1 pid(进程号)，所以，安装了虚拟机——Linux系统：  
开源镜像站：https://mirrors.aliyun.com/centos/7/isos/x86_64/?spm=a2c6h.25603864.0.0.685845113SlugF  
Centos7 安装教程：https://blog.csdn.net/u014644574/article/details/112494541
* 借助 FinalShell SSH工具把编译好的可执行的程序发送到虚拟机中的 Linux 环境  
编译的时候需要修改 ENV ：`go env -w GOOS=linux`，`go env -w GOARCH=adm64`  
FinalShell SSH工具 Windows 版本下载：http://www.hostbuf.com/downloads/finalshell_install.exe

## 项目依赖
1. 搭建项目
* `go`
* `go get -u github.com/gin-gonic/gin`  
* `go get -u github.com/swaggo/swag/cmd/swag`
* `go get -u github.com/swaggo/gin-swagger`
* `go get -u github.com/swaggo/files`

2. 配置文件
* `go get -u gopkg.in/ini.v1`

3. 记录日志
* 标准库：`fmt` `log` `os` `time`
* 标准库： `path/filepath` `runtime`

4. 优雅地关机
* 标准库：`net/http` `os` `os/signal` `syscall`

5. 优雅地重启
* `go get -u github.com/fvbock/endless`

## 项目结构
- conf : 用于存储配置文件
- docs : Swag 自动生成的 API 文档 
- pkg : 用于存储项目自定义的包
- routers : 用于处理路由（处理路由的逻辑代码）
- main : 项目入口