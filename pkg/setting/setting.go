// setting 包实现了读取 ini 配置文件
package setting

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File // ini 数据源

	// 应用的模式
	RunMode string //【debug 开发 release 发布】

	// server 配置
	HTTPHost     string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Static       string

	// database 配置
	DbType        string
	DbUser        string
	DbPassword    string
	DbHost        string
	DbName        string
	DbTablePrefix string

	// log
	LogSavePath string // 日志保存路径
	LogSaveName string // 日志文件保存名称（log+时间）
	LogFileExt  string // 日志文件类型
	TimeFormat  string // 时间格式
)

// Init 用于初始化配置文件
func init() {
	var err error
	// 加载并解析INI数据源
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println("加载解析 conf/app.ini 错误", err)
		return
	}
	// 加载应用的模式【debug 开发 release 发布】
	loadBase()

	// loadServer 加载 server 服务器配置
	loadServer()

	// laodDatabase 加载配置文件中的 Database 信息
	loadDatabase()

	// laodLog 加载配置文件中的 log 信息
	loadLog()

}

// loadBase 加载配置中的基础信息【应用的模式】
func loadBase() {
	// Section("") 表示是默认分区
	// Key("RUN_MODE") 表示操作键
	// MustString("debug") 当操作键不存在或者转换失败时，使用默认值 debug
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// loadServer 加载 server 服务器配置
func loadServer() {
	// HTTP_PORT = 8000
	// READ_TIMEOUT = 60
	// WRITE_TIMEOUT = 60
	// Static = public
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("获取 server 分区失败：%v", err)
	}
	HTTPHost = sec.Key("HTTP_HOST").MustString("0.0.0.0")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	Static = sec.Key("STATIC").MustString("public")
}

// loadDatabase 加载配置文件中的 Database 信息
func loadDatabase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("获取 database 分区失败：%v", err)
	}
	DbType = sec.Key("DB_TYPE").MustString("mysql")
	DbUser = sec.Key("DB_USER").MustString("root")
	DbPassword = sec.Key("DB_PASSWORD").MustString("Aa_123456")
	DbHost = sec.Key("DB_HOST").MustString("127.0.0.1:3306")
	DbName = sec.Key("DB_NAME").MustString("chitchat")
	DbTablePrefix = sec.Key("DB_TABLE_PREFIX").MustString("cc_")
}

// loadLog 加载配置文件中的 log 信息
func loadLog() {
	// LogSavePath 日志保存路径
	LogSavePath = Cfg.Section("log").Key("LOG_SAVE_PATH").MustString("runtime/logs")
	// LogSaveName 日志文件保存名称
	LogSaveName = Cfg.Section("log").Key("LOG_SAVE_NAME").MustString("log")
	// LogFileExt 日志文件的后缀名
	LogFileExt = Cfg.Section("log").Key("LOG_FILE_EXT").MustString("log")
	// TimeFormat 日志保存时间的格式
	TimeFormat = Cfg.Section("log").Key("TIME_FORMAT").MustString("20060102")
}
