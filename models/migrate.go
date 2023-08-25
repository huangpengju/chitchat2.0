package models

func migration() {
	// 创建表时添加表后缀信息：数据库引擎、字符集
	db.Set("gorm:table_options", "ENGINE=InnoDB charset=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(&User{}).AutoMigrate(&UserInfo{})
	// 给 UserInfo 表，设置外键
	db.Model(&UserInfo{}).AddForeignKey("user_id", "cc_user(id)", "CASCADE", "CASCADE")
}
