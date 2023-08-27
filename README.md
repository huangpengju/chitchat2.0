# ChitChat2.0 论坛
ChitChat2.0 仓库是中 ChitChat 项目的升级版 

## 项目实现
后端使用 Go + Gin + Gorm + JWT + Mysql  
前端使用 HTML5 + Vue
## 项目思路
1. 搭建项目
* Go + Gin 开始搭建项目
* 为已搭建的项目加上 Swagger，Swag 将 Go 的注释转换为 Swagger2.0 文档  
https://github.com/swaggo/swag/blob/master/README_zh-CN.md
* Swag 支持 Gin 框架，实现 Swag 与 Gin 集成  
https://github.com/swaggo/gin-swagger

2. 配置文件
* 创建 ini 文件。
* 读取 ini 中的内容信息，应用到 gin.SetMode 和记录日志。

3. 记录日志
* 日志目录和日志文件：存在则打开日志文件，不存在则创建。
* 日志的记录：初始化时先打开日志文件，调用日志记录函数时设置日志前缀，输出日志。

4. 优雅地关机
* http.Server - Shutdown()  
优雅关机就是服务端关机命令发出后不是立即关机，而是等待当前还在处理的请求全部处理完毕后再退出程序，是一种对客户端友好的关机方式。而执行Ctrl+C关闭服务端时，会强制结束进程导致正在访问的请求出现问题。

5. 优雅地重启
* 使用 fvbock/endless 来替换默认的 ListenAndServe 启动服务来实现。  
由于开发环境是 Windows ，且参考资料中需要使用命令： kill -1 pid(进程号)，关闭进程，所以，安装了虚拟机——Linux系统：  
开源镜像站：https://mirrors.aliyun.com/centos/7/isos/x86_64/?spm=a2c6h.25603864.0.0.685845113SlugF  
Centos7 安装教程：https://blog.csdn.net/u014644574/article/details/112494541
借助 FinalShell SSH工具把编译好的可执行的程序发送到虚拟机中的 Linux 环境  
编译的时候需要修改 ENV ：`go env -w GOOS=linux`，`go env -w GOARCH=amd64`，`go build main.go` 或者 `set GOARCH=amd64`，`set GOOS=linux`，`go build main.go`  
FinalShell SSH工具 Windows 版本下载：http://www.hostbuf.com/downloads/finalshell_install.exe

6. golang 程序的热加载
* 所谓热加载就是当对代码进行修改时，程序能够自动重新加载并执行，这在开发中是非常便利的，可以快速进行代码测试，省去了每次手动重新编译。  
beego 中可以使用官方给提供bee工具来热加载项目，但是 gin 中并没有官方提供的热加载工具，这个时候要实现热加载就可以借助第三方的工具。  
工具 1（推荐）：https://github.com/gravityblast/fresh  
工具 2：https://github.com/codegangsta/gin

7. HTML 模板渲染
* 参考 gin Web Framework :https://gin-gonic.com/zh-cn/docs/examples/html-rendering/  
其他博客：https://blog.csdn.net/zhoupenghui168/article/details/128996683  
先设置一个路由，并渲染出该路由的模板：`layout.html` ，然后给`layout.html`模板加载`css` `js`资源,再引入其他嵌套模板到`layout.html`中。最后设置其他路由，实现页面跳转。

8. 连接数据库
* 创建建库，设计数据库表；
* 下载 gorm 依赖和 mysql 数据库驱动依赖；
* 读取数据库配置信息后连接 mysql 数据库,设置表前缀和不加 s,启用日志和设置连接池；
* 表自动迁移同时设置表关系（外键关联）和 表后缀信息（ENGINE=InnoDB charset=utf8mb4 COLLATE=utf8mb4_unicode_ci）

9. 实现注册、实现登录
* 声明用于接收注册数据的 userSerive(用户服务) 结构体，使用 c.ShouldBind 进行数据绑定，同时进行表单验证。  
表单验证提示错误消息是英语，不友好，对错误消息进行翻译。
* 准备一个 serializer(序列化数据) 结构体，用于注册成功时，封装返回的数据。  
封装响应状态码（Status）、数据（Data）、消息（Message）、错误(Error)
* 开始实现用户注册  
根据账号查询用户是否存在、对密码进行加密处理、开启事物、注册用户、提交事物、返回封装的数据
*  根据返回数据中的响应状态码（Status）跳转页面
* 开始实现用户登录  
根账号查询用户是否存在、验证登录用户密码、生成Token、返回封装的数据
*  根据返回数据中的响应状态码（Status）跳转页面


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

6. golang 程序的热加载
* 工具1：`go install github.com/pilu/fresh@latest`，`go get github.com/pilu/fresh` 然后运行命令 fresh

* 工具2：`go get -u github.com/codegangsta/gin` 然后运行命令 `gin run main.go`` 

7. HTML 模板渲染
* r.LoadHTMLGlob

8. gorm 和 驱动程序
* `go get -u gorm.io/gorm` 或 `go get -u github.com/jinzhu/gorm` 国内一般用后者
* `go get -u github.com/jinzhu/gorm/dialects/mysql`

9. Token
* `go get -u github.com/dgrijalva/jwt-go`

## 项目结构
- conf : 用于存储配置文件
- docs : Swag 自动生成的 API 文档 
- pkg : 用于存储项目自定义的包
- routers : 用于处理路由（处理路由的逻辑代码）
- main : 项目入口