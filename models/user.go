package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50);not null;unique;comment:'用户名'"`
	Password string `gorm:"type:varchar(100);not null;comment:'密码'"`
	Email    string `gorm:"type:varchar(100);comment:'邮箱'"`
	Phone    string `gorm:"type:varchar(20);comment:'电话'"`
}
