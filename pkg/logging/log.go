// logging 包中 log.go 实现了日志文件的记录
package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File // 文件对象

	DefaultPrefix      = "" // 日志的默认前缀（每个生成的日志行的开头，赋值为"" 相当于日志不设置前缀）
	DefaultCallerDepth = 2  //  默认调用方深度（用于 runtime.Caller 获取go程序调用栈所执行的函数的文件和行号）

	logger     *log.Logger                                           // 日志对象
	logPrefix  = ""                                                  // 给日志设置特定格式的前缀
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"} // 日志级别标签
)

// 定义常量，当作键，用于取出 levelFlags 切片的值
const (
	DEBUG   Level = iota // 0
	INFO                 // 1
	WARNING              // 2
	ERROR                // 3
	FATAL                // 4
)

func init() {
	// 获取日志路径：runtime/logs/log20230822.log
	filePath := getLogFileFullPath()

	// openLogFile 先检查日志目录和日志文件，不存在则创建，返回日志文件对象
	F = openLogFile(filePath)

	// New创建一个 Logger （日志对象）
	// 参数1 设置日志信息写入的目的地（定义要写入日志数据的IO句柄）【F: *os.File】。
	// 参数2 会添加到生成的每一条日志前面（定义每个生成的日志行的开头）【DefaultPrefix：""】。
	// 参数3 定义日志的记录属性（时间、文件等等）【log.LstdFlags ：日志记录的格式属性之一|2009/01/23 01:23:23 】。
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// setPrefix 自定义设置每条日志的前缀
// 参数 level 是日志级别【常量：DEBUG=0、INFO=1、WARNING=2、ERROR=3、FATAL=4】
func setPrefix(level Level) {
	// 思路：获取抛出日志的位置（行号）、以及所在文件，拼接相应的字符后使用 SetPrefix 进行设置
	// runtime.Caller(0) 时，setPrefix → runtime.Caller(0),此时 调用者函数是 setPrefix ，file应该是 setPrefix所在文件
	// runtime.Caller(1) 时，A → setPrefix → runtime.Caller(1) ,此时调用者函数是 A，file 应该是A 所在文件
	// runtime.Caller(2) 时，B → A → SetPrefix → runtime.Caller(2),此时调用者函数是 B，file 应该是B所在文件
	_, file, line, ok := runtime.Caller(DefaultCallerDepth) // DefaultCallerDepth 表示要找的函数调用者深度，深度为0表示setPrefix，
	if ok {
		// filepath.Base() 返回路径的最后一个元素，也就是 xxx.go
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	// SetPrefix 设置标准 logger 日志对象的输出前缀。
	logger.SetPrefix(logPrefix)
}

// Debug 日志记录器
func Debug(v ...interface{}) {
	// 设置日志的前缀
	setPrefix(DEBUG)
	// 输出日志
	logger.Println(v...)
}

// Info 日志记录器
func Info(v ...interface{}) {
	// 设置日志的前缀
	setPrefix(INFO)
	// 输出日志
	logger.Println(v...)
}

// Warn 日志记录器
func Warn(v ...interface{}) {
	// 设置日志的前缀
	setPrefix(WARNING)
	// 输出日志
	logger.Println(v...)
}

// Error 日志记录器
func Error(v ...interface{}) {
	// 设置日志的前缀
	setPrefix(ERROR)
	// 输出日志
	logger.Println(v...)
}

// Fatal 日志记录器
func Fatal(v ...interface{}) {
	// 设置日志的前缀
	setPrefix(FATAL)
	// 输出日志
	logger.Println(v...)
}
