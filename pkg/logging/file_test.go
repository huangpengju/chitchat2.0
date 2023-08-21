package logging

import (
	"fmt"
	"testing"

	"chitchat2.0/pkg/setting"
)

func TestGetLogFilePullPath(t *testing.T) {
	setting.Init("C:/Users/Vcom/Desktop/src/chitchat2.0/conf/app.ini")

	path := getLogFileFullPath()
	if path == "" {
		t.Error("错误的内容，期待是 /runtime/logs/log20230821.log,但获取到了", path)
	}
	fmt.Println("path=", path)
}
