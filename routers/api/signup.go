package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "default/login.layout.html", gin.H{
		"state": 0,
	})
}

func SignupAccount(c *gin.Context) {
	c.String(http.StatusOK, "注册")
}
