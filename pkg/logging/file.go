// logging 包实现了日志文件的生成以及记录日志
package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"chitchat2.0/pkg/setting"
)

// getLogFileFullPath 获取日志文件的完成路径(相对路径+文件)
func getLogFileFullPath() string {
	// 相对路径
	prefixPath := setting.LogSavePath
	// 文件名（log）+ 时间 + 文件后缀名
	suffixPath := fmt.Sprintf("%s%s.%s", setting.LogSaveName, time.Now().Format(setting.TimeFormat), setting.LogFileExt)

	// 相对路径 + 文件 （/runtime/logs/log20230821.log）
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// openLogFile 根据完整的 log 路径打开 log 文件
// log 文件不存在则创建
func openLogFile(filePath string) *os.File {
	// 返回指定文件对象
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err): // 返回一个布尔值说明该错误是否表示一个文件或目录不存在。
		// 创建文件
		mkDir()
	case os.IsPermission(err): //返回一个布尔值说明该错误是否表示因权限不足要求被拒绝。
		log.Fatalf(":", err)
	}
}

func mkDir() {
	// os.Getwd :返回当前目录对应的根目录
	dir, _ := os.Getwd()

}
