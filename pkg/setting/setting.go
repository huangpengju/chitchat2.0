// setting 包实现了读取 ini 配置文件
package setting

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File // ini 数据源

	// 应用的模式
	RunMode string //【debug 开发 release 发布】

	// log
	LogSavePath string // 日志保存路径
	LogSaveName string // 日志文件保存名称（log+时间）
	LogFileExt  string // 日志文件类型
	TimeFormat  string // 时间格式
)

// Init 用于初始化配置文件
func Init(iniPath string) {
	fmt.Println("iniPath==", iniPath)
	var err error
	// 加载并解析INI数据源
	// Cfg, err = ini.Load("conf/app.ini")

	Cfg, err = ini.Load(iniPath)
	if err != nil {
		fmt.Println("加载解析 conf/app.ini 错误", err)
		return
	}
	// 加载应用的模式【debug 开发 release 发布】
	loadBase()

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

// laodLog 加载配置文件中的 log 信息
func loadLog() {
	// LogSavePath 日志保存路径
	LogSavePath = Cfg.Section("log").Key("LOG_SAVE_PATH").MustString("/runtime/logs")
	// LogSaveName 日志文件保存名称
	LogSaveName = Cfg.Section("log").Key("LOG_SAVE_NAME").MustString("log")
	// LogFileExt 日志文件的后缀名
	LogFileExt = Cfg.Section("log").Key("LOG_FILE_EXT").MustString("log")
	// TimeFormat 日志保存时间的格式
	TimeFormat = Cfg.Section("log").Key("TIME_FORMAT").MustString("20060102")
}
