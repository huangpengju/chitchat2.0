package api

import (
	"net/http"

	"chitchat2.0/pkg/httputil"
	"chitchat2.0/service"
	"github.com/gin-gonic/gin"
)

// Signup 加载注册页面
func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
		"state": 0,
	})
}

// SignupAccount 用户注册接口文档
// @Summary 用户注册
// @Description SignupAccount 接口实现了用户注册功能
// @Tags User
// @Accept json
// @produce json
// @Param username query string true "UserName(账户)"
// @Param email query string true "Email(邮箱)"
// @Param password query string true "Password(密码)"
// @Success 200 {object} serializer.Response "注册成功"
// @Failure 400 {object} httputil.HTTPError "参数验证失败"
// @Router /signup_account [post]
func SignupAccount(c *gin.Context) {
	// c.String(http.StatusOK, "注册")
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		// 绑定成功后，进行注册
		// serializer := userService.Register()
		serializer := userService.RegisterBegin()
		c.JSON(http.StatusOK, serializer)
	} else {
		httputil.NewError(c, http.StatusBadRequest, err)
	}

}
