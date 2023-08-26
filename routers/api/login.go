package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 加载登录页面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
		"state": 1,
	})
}

// Authenticate 用户登录接口文档
// @Summary 用户登录
// @Description Authenticate 接口实现了用户登录功能
// @Tags User
// @Accept json
// @Param username query string true "UserName(账户)"
// @Param password query string true "Password(密码)"
// @Success 200 {object} service.UserService
// @Failure 400 {string} string "参数绑定失败"
// @Router /authenticate [post]
func Authenticate(c *gin.Context) {
	c.String(http.StatusOK, "开始登录")

}
