package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50);not null;unique;comment:'用户名'"`
	Password string `gorm:"type:varchar(100);not null;comment:'密码'"`
	Email    string `gorm:"type:varchar(100);comment:'邮箱'"`
	Phone    string `gorm:"type:varchar(20);comment:'电话'"`
}

// SetPassword 实现加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
