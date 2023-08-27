package api

import (
	"net/http"

	"chitchat2.0/pkg/validator"
	"chitchat2.0/service"
	"github.com/gin-gonic/gin"
)

// Login 加载登录页面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
		"state": 1,
		"msg":   "",
	})
}

// Authenticate 用户登录接口文档
// @Summary 用户登录
// @Description Authenticate 接口实现了用户登录功能
// @Tags User
// @Accept json
// @produce json
// @Param username query string true "UserName(账户)"
// @Param password query string true "Password(密码)"
// @Success 200 {object} serializer.Response "登录成功"
// @Failure 400 {object} httputil.HTTPError "登录失败"
// @Router /authenticate [post]
func Authenticate(c *gin.Context) {
	var userLogin service.UserLoginService

	if err := c.ShouldBind(&userLogin); err == nil {
		serializer := userLogin.Login()
		if serializer.Status != 200 {
			c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
				"state": 1,
				"msg":   serializer.Message,
			})
		} else {
			users := []UserInfo{
				{Uuid: 1, Name: "李四1", Gender: "男", Age: 17, CreatedAtDate: "2023-08-22", NumReplies: 1, Topic: "前台1号---张三"},
				{Uuid: 2, Name: "王二2", Gender: "男", Age: 18, CreatedAtDate: "2023-08-23", NumReplies: 2, Topic: "前台2号---王二"},
				{Uuid: 3, Name: "麻子3", Gender: "男", Age: 19, CreatedAtDate: "2023-08-24", NumReplies: 3, Topic: "前台3号---麻子"}}
			c.HTML(http.StatusOK, "default/layout.html", gin.H{
				"state": 0,
				"users": users,
			})
		}

	} else {
		c.HTML(http.StatusBadRequest, "default/login.layout.html", gin.H{
			"state":   1,
			"message": validator.Translate(err),
		})
	}
}

type UserInfo struct {
	Name          string
	Gender        string
	Age           int
	CreatedAtDate string
	NumReplies    int
	Topic         string
	Uuid          int
}
