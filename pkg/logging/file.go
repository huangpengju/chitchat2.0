// logging 包中 file.go 实现了日志文件及文件目录的创建
package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"chitchat2.0/pkg/setting"
)

// getLogFileFullPath 获取日志文件的完成路径(路径+文件：runtime/logs/log20230822.log)
func getLogFileFullPath() string {
	// 相对路径 runtime/logs/
	prefixPath := setting.LogSavePath
	// setting.LogSaveName log
	// TimeFormat 20230822
	// setting.LogFileExt log
	// 文件名（log）+ 时间 + 文件后缀名   log20230822.log
	suffixPath := fmt.Sprintf("%s%s.%s", setting.LogSaveName, time.Now().Format(setting.TimeFormat), setting.LogFileExt)

	// 相对路径 + 文件 （runtime/logs/log20230822.log）
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// openLogFile 先检查日志目录和日志文件，不存在则创建，返回日志文件对象
func openLogFile(filePath string) *os.File {
	//1. 先根据完整的 runtime/logs/log20230822.log 判断目录或者文件是否存在，不存在则创建
	_, err := os.Stat(filePath) // 返回文件信息结构描述文件。如果出现错误，会返回 *PathError
	switch {
	case os.IsNotExist(err): // 返回一个布尔值说明该错误是否表示一个文件或目录不存在。
		// 创建目录
		mkDir()
	case os.IsPermission(err): //返回一个布尔值说明该错误是否表示因权限不足要求被拒绝。
		log.Fatalf("日志目录或者文件的权限问题:%v", err)
	}

	// 2.然后，根据完整的 /runtime/logs/log20230822.log 打开文件，不存在则创建
	// os.OpenFile：调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于 I/O。如果出现错误，则为*PathError。
	// os.O_APPEND 在写入数据时追加到文件中。
	// os.O_CREATE 如果不存在，请创建一个新文件。
	// os.O_WRONLY 以只写模式打开文件
	// 0644->即用户具有读写权限，组用户和其它用户具有只读权限；
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开日志文件失败：%v", err)
	}
	return handle
}

// mkDir 创建日志存放目录：chitchat2.0/runtime/logs/
func mkDir() {
	// os.Getwd :返回当前目录对应的根目录
	dir, _ := os.Getwd()
	// os.MkdirAll 创建对应的目录以及所需的子目录，若成功则返回 nil,否则返回err
	// dir 项目运行目录（测试时是测试目录），setting.LogSavePath 是 /runtime/logs/
	// ModePerm FileMode = 0777  0777->即用户，组用户和其它用户具有读/写/执行权限；
	err := os.MkdirAll(dir+"/"+setting.LogSavePath, os.ModePerm)
	if err != nil {
		// "创建日志目录出现错误",
		panic(err)
	}

}
