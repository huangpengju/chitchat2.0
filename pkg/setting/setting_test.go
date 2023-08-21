package setting

import (
	"fmt"
	"testing"
)

func TestLoadLog(t *testing.T) {

	Init("C:/Users/Vcom/Desktop/src/chitchat2.0/conf/app.ini")
	fmt.Println("RunMode=", RunMode)
	fmt.Println("LogSavePath=", LogSavePath)
	fmt.Println("LogSaveName=", LogSaveName)
	fmt.Println("LogFileExt=", LogFileExt)
	fmt.Println("TimeFormat=", TimeFormat)
}
