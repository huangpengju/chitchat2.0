package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"chitchat2.0/pkg/setting"
)

func TestGetLogFilePullPath(t *testing.T) {
	path, _ := os.Getwd()
	fmt.Println("path==", path)
	if strings.HasSuffix(path, "logging") {
		path1 := filepath.Dir(path)
		fmt.Println("path1==", path1)
		if strings.HasSuffix(path1, "pkg") {
			path2 := filepath.Dir(path1)
			_ = os.Chdir(path2)
			setting.Init(path2)
		}
	}
	// 测试
	path3 := getLogFileFullPath()
	if path3 == "" {
		t.Error("错误的内容，期待是 /runtime/logs/log20230821.log,但获取到了", path)
	}
	fmt.Println("path3=", path3)
}
