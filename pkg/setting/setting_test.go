package setting

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestLoadLog(t *testing.T) {
	var err error
	// 加载并解析INI数据源
	Cfg, err = ini.Load("conf/app.ini")

	// Cfg, err = ini.Load(iniPath)
	if err != nil {
		fmt.Println("加载解析 conf/app.ini 错误", err)
		return
	}
	// 加载应用的模式【debug 开发 release 发布】
	loadBase()

	// laodLog 加载配置文件中的 log 信息
	loadLog()

	fmt.Println("RunMode=", RunMode)
	fmt.Println("LogSavePath=", LogSavePath)
	fmt.Println("LogSaveName=", LogSaveName)
	fmt.Println("LogFileExt=", LogFileExt)
	fmt.Println("TimeFormat=", TimeFormat)
}
