package models

import "github.com/jinzhu/gorm"

type UserInfo struct {
	gorm.Model

	// User User `gorm:"ForeignKey:UserID"` //标签设置外键无效，设置外键在表迁移之后设置
	User    User
	UserID  uint   `gorm:"not null;comment:'用户ID'"`
	Name    string `gorm:"type:varchar(50);comment:'姓名'"`
	Age     int    `gorm:"type:tinyint(3);comment:'年龄'"`
	Gender  int    `gorm:"type:tinyint(3);default:'0';comment:'0-位置，1-男，2-女'"`
	Avatar  string `gorm:"type:varchar(100);comment:'头像'"`
	Address string `gorm:"type:varchar(100);comment:'地址'"`
	Intro   string `gorm:"type:varchar(500);comment:'简介'"`
}
