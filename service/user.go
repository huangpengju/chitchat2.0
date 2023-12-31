package service

import (
	"fmt"
	"net/http"

	"chitchat2.0/models"
	"chitchat2.0/pkg/serializer"
	"chitchat2.0/pkg/utils"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"username" binding:"required,min=3,max=15"`
	Password string `form:"password" binding:"required,min=5,max=16"`
	Email    string `form:"email" binding:"email"`
}

type UserLoginService struct {
	UserName string `form:"username" binding:"required,min=3,max=15"`
	Password string `form:"password" binding:"required,min=5,max=16"`
}

// Register 实现用户注册(未开事物)
func (userService *UserService) Register() serializer.Response {
	// 1.先检查用户是否存在
	// 2.给模型赋值（对密码进行加密）
	// 3.实现注册

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
			Message: err.Error(),           // 密码加密错误
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

// RegisterBegin 用户注册并开启事物
func (userService *UserService) RegisterBegin() serializer.Response {
	var user models.User
	var userInfo models.UserInfo

	var count int
	models.Db.Model(&models.User{}).Where("user_name=?", userService.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status:  http.StatusBadRequest, // 400
			Message: fmt.Sprintf("账号“%v”已存在", userService.UserName),
		}
	}
	user.UserName = userService.UserName
	user.Email = userService.Email
	if err := user.SetPassword(userService.Password); err != nil {
		return serializer.Response{
			Status:  http.StatusBadRequest, //400
			Message: err.Error(),           // 密码加密错误
		}
	}
	// 开始注册用户
	// 开始事物
	tx := models.Db.Begin()

	// 执行数据库操作
	if err := tx.Create(&user).Error; err != nil {
		// 出现错误则回滚事物
		tx.Rollback()
		return serializer.Response{
			Status:  http.StatusInternalServerError, //500
			Message: "数据库操作错误",
		}
	}
	userInfo.UserID = user.ID
	if err := tx.Create(&userInfo).Error; err != nil {
		// 出现错误则回滚事物
		tx.Rollback()
		return serializer.Response{
			Status:  http.StatusInternalServerError, //500
			Message: "数据库操作错误",
		}
	}
	// 提交事物
	if err := tx.Commit().Error; err != nil {
		// 出现错误则回滚事物
		return serializer.Response{
			Status:  http.StatusInternalServerError, //500
			Message: "数据库操作错误",
		}
	}
	return serializer.Response{
		Status:  http.StatusOK, // 200
		Message: "账号注册成功",
	}
}

// Login 用户登录
func (userService *UserLoginService) Login() serializer.Response {

	var user models.User
	// First 检索单个对象，没有找到记录时，它会返回 ErrRecordNotFound 错误
	if err := models.Db.Where("user_name=?", userService.UserName).First(&user).Error; err != nil {
		//
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status:  http.StatusBadRequest, //400
				Data:    err,
				Message: "账户不存在，请先注册",
			}
		}
		// 用户存在，其他因素的错误
		return serializer.Response{
			Status:  http.StatusInternalServerError, //500
			Data:    err,
			Message: "数据库查询用户信息出错",
		}
	}
	// 找到用户
	// 去验证登录用户的密码
	if !user.CheckPassword(userService.Password) {
		// 密码不相同
		return serializer.Response{
			Status:  http.StatusBadRequest, //400
			Message: "密码错误",
		}
	}

	// 准备一个 token,作为响应返回
	token, err := utils.GenerateToken(user.ID, user.UserName, userService.Password)
	if err != nil {
		return serializer.Response{
			Status:  http.StatusInternalServerError, //500
			Message: "Token 签发出错",
		}
	}

	// 返回用户信息
	return serializer.Response{
		Status:  http.StatusOK,
		Data:    serializer.TokenData{User: user, Token: token},
		Message: "登录成功",
	}
}
