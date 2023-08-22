package setting

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadLog(t *testing.T) {
	// Getwd返回一个对应当前工作目录的根路径。
	path, _ := os.Getwd()
	// 判断 path2 是否有后缀字符串 setting
	if strings.HasSuffix(path, "setting") {
		// Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录
		path1 := filepath.Dir(path)

		if strings.HasSuffix(path1, "pkg") {
			path2 := filepath.Dir(path1)
			_ = os.Chdir(path2)
			// 测试
			// Init(path2)
		}
	}
	fmt.Println("RunMode=", RunMode)
	fmt.Println("LogSavePath=", LogSavePath)
	fmt.Println("LogSaveName=", LogSaveName)
	fmt.Println("LogFileExt=", LogFileExt)
	fmt.Println("TimeFormat=", TimeFormat)
}
