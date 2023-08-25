package models

import (
	"fmt"
	"log"
	"strings"

	"chitchat2.0/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	DSN := strings.Join([]string{
		setting.DbUser,
		":",
		setting.DbPassword,
		"@tcp(", setting.DbHost, ")/",
		setting.DbName,
		"?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	db, err = gorm.Open(setting.DbType, DSN)
	if err != nil {
		log.Panic("数据库连接失败==", err)
		return
	}
	fmt.Println("======数据库连接成功")

	// 更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		// 设置指定的表前缀
		return setting.DbTablePrefix + defaultTableName
	}

	// 设置表名不加 s
	db.SingularTable(true)
	// 启用日志
	db.LogMode(true)
	// 设置连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// 自动迁移表
	migration()
}
