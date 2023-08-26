package httputil

import (
	"chitchat2.0/pkg/validator"
	"github.com/gin-gonic/gin"
)

// 表单验证失败，响应消息设置
// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: validator.Translate(err),
	}
	// Message: err.Error(),
	ctx.JSON(status, er)
}

// 表单验证，响应结构
// HTTPError example
type HTTPError struct {
	Code int `json:"code" example:"400"`
	// Message string `json:"message" example:"status bad request"`
	Message map[string][]string `json:"message" `
}
