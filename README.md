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

3. 记录日志
* 读取配置文件，获取日志配置信息
* 根据日志路径（拼接而成）访问日志，若不存在则创建（文件夹和文件）

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

## 项目结构
- conf : 用于存储配置文件
- docs : Swag 自动生成的 API 文档 
- pkg : 用于存储项目自定义的包
- routers : 用于处理路由（处理路由的逻辑代码）
- main : 项目入口