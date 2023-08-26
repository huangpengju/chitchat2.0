package api

import (
	"net/http"

	"chitchat2.0/pkg/validator"
	"chitchat2.0/service"
	"github.com/gin-gonic/gin"
)

// Signup 加载注册页面
func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
		"state": 0,
		"msg":   "在下面注册一个账户",
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
		// c.Redirect(http.StatusFound, "/login")
		if serializer.Status != 200 {
			c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
				"state": 0,
				"msg":   serializer.Message,
			})
		} else {
			c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
				"state": 1,
				"msg":   serializer.Message,
			})
		}
	} else {
		// httputil.NewError(c, http.StatusBadRequest, err)

		c.HTML(http.StatusBadRequest, "default/login.layout.html", gin.H{
			"state":   0,
			"message": validator.Translate(err),
		})
	}

}
