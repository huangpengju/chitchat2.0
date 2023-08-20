# ChitChat2.0 论坛
ChitChat2.0 仓库是中 ChitChat 项目的升级版 

## 项目实现
后端使用 Go + Gin + Gorm + JWT + Mysql  
前端使用 HTML5 + Vue

## 项目依赖
实现 gin 项目搭建
* go     `安装 Go`
* gin    `go get -u github.com/gin-gonic/gin`  

为项目加上 Swagger，Swag 将 Go 的注释转换为Swagger2.0文档
* go get -u github.com/swaggo/swag/cmd/swag

Swag 支持 Gin 框架，下面的依赖实现 Swag 与 Gin 集成
* go get -u github.com/swaggo/gin-swagger
* go get -u github.com/swaggo/files