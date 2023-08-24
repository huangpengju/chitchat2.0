package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
		"state": 1,
	})
}

func Authenticate(c *gin.Context) {
	c.String(http.StatusOK, "开始登录")

}
