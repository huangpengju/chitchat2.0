package service

import (
	"fmt"
	"net/http"

	"chitchat2.0/models"
	"chitchat2.0/pkg/serializer"
)

type UserService struct {
	UserName string `form:"username" binding:"required,min=3,max=15"`
	Password string `form:"password" binding:"required,min=5,max=16"`
	Email    string `form:"email" binding:"email"`
}

// Register 实现用户注册
func (userService *UserService) Register() serializer.Response {
	// 1.先检查用户是否存在
	// 2.给模型赋值（对密码进行加密）
	// 3.实现注册
	fmt.Println(userService)
	// 数据库 User 模型 和 UserInfo 模型
	var user models.User
	var userInfo models.UserInfo
	// 用户个数
	var count int

	// 判断用户是否存在
	models.Db.Model(&models.User{}).Where("user_name=?", userService.UserName).First(&user).Count(&count)
	if count == 1 {
		// 用户名已经存在，返回错误消息
		msg := fmt.Sprintf("用户“%v”已存在", userService.UserName)
		return serializer.Response{
			Status:  http.StatusBadRequest, //400
			Message: msg,
		}
	}
	// 用户名验证通过后，把接收到的数据赋值给数据库 user 模型
	user.UserName = userService.UserName
	user.Email = userService.Email
	// 对密码加密,同时把密码赋值给 user 模型
	if err := user.SetPassword(userService.Password); err != nil {
		return serializer.Response{
			Status:  http.StatusBadRequest, //400
			Message: err.Error(),
		}
	}
	// 开始注册用户
	if err := models.Db.Create(&user).Error; err != nil {
		return serializer.Response{
			Status:  http.StatusInternalServerError, // 500
			Message: "数据库操作错误",
		}
	}
	userInfo.UserID = user.ID
	if err := models.Db.Create(&userInfo).Error; err != nil {
		return serializer.Response{
			Status:  http.StatusInternalServerError, // 500
			Message: "数据库操作错误",
		}
	}
	return serializer.Response{
		Status:  http.StatusOK, // 200
		Message: "用户注册成功",
	}
}
