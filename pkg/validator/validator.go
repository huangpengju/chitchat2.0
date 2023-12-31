// validator 修改表单验证提示消息为中文显示
package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

// 返回 service 中参数 binding 验证消息
func init() {
	// 注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	// 获取gin的效验器
	validate := binding.Validator.Engine().(*validator.Validate)
	// 注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

}

// Translate 翻译错误信息
func Translate(err error) map[string][]string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}
