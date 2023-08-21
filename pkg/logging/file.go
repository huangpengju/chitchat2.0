// logging 包实现了日志文件的生成以及记录日志
package logging

import (
	"fmt"
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

//
