package api

import "github.com/gin-gonic/gin"

type auth struct {
	Username string
	Password string
}

func GetAuth(c *gin.Context) {

}
