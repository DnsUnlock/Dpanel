package log

import (
	"bytes"
	"fmt"
	"github.com/lcvvvv/stdio"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Level int

const (
	DEBUG Level = 0x0000a1
	INFO        = 0x0000b2
	WARN        = 0x0000c3
	ERROR       = 0x0000d4
	DATA        = 0x0000f5
	NONE        = 0x0000e6
)

type Logger struct {
	log      *log.Logger
	modifier func(string) string
	filter   func(string) bool
}

func (l *Logger) Printf(format string, s ...interface{}) {
	expr := fmt.Sprintf(format, s...)
	l.Println(expr)
}

func (l *Logger) Println(s ...interface{}) {
	expr := fmt.Sprint(s...)
	if l.modifier != nil {
		expr = l.modifier(expr)
	}
	if l.filter != nil {
		if l.filter(expr) == true {
			return
		}
	}
	l.log.Output(3, expr)
}

var info = &Logger{
	log.New(stdio.Out, "\r[+]", log.Ldate|log.Ltime|log.Llongfile),
	Green,
	nil,
}

var warn = &Logger{
	log.New(stdio.Out, "\r[*]", log.Ldate|log.Ltime|log.Llongfile),
	Red,
	nil,
}

var err = &Logger{
	log.New(stdio.Out, "\r[Error]", log.Ldate|log.Ltime|log.Llongfile),
	Red,
	nil,
}

var dbg = &Logger{
	log.New(stdio.Out, "\r[-]", log.Ldate|log.Ltime|log.Llongfile),
	debugModifier,
	debugFilter,
}

func debugModifier(s string) string {
	_, file, line, _ := runtime.Caller(3)
	file = file[strings.LastIndex(file, "/")+1:]
	logStr := fmt.Sprintf("%s%s(%d) %s", "> ", file, line, s)
	logStr = Yellow(logStr)
	return logStr
}

func debugFilter(_ string) bool {
	//Debug 过滤器
	//if strings.Contains(s, "STEP1:CONNECT") {
	//	return true
	//}
	return false
}

var data = &Logger{
	log.New(stdio.Out, "\r", 0),
	nil,
	nil,
}

func Printf(level Level, format string, s ...interface{}) {
	Println(level, fmt.Sprintf(format, s...))
}

func Println(level Level, s ...interface{}) {
	logStr := fmt.Sprint(s...)
	switch level {
	case DEBUG:
		dbg.Println(logStr)
	case INFO:
		info.Println(logStr)
	case WARN:
		warn.Println(logStr)
	case ERROR:
		err.Println(logStr)
		os.Exit(0)
	case DATA:
		data.Println(logStr)
	default:
		return
	}
}

var empty = &Logger{log.New(io.Discard, "", 0), nil, nil}

func SetLevel(level Level) {
	if level > ERROR {
		err = empty
	}
	if level > WARN {
		warn = empty
	}
	if level > INFO {
		info = empty
	}
	if level > DEBUG {
		dbg = empty
	}
	if level > NONE {
		//nothing
	}
}

func SetOutput(writer io.Writer) {
	data.modifier = func(s string) string {
		_, _ = writer.Write([]byte(Clear(s)))
		_, _ = writer.Write([]byte("\r\n"))
		return s
	}
}

func SetOutputFile(level Level, fileName string) {
	// 获取文件路径中的目录部分
	dirName := filepath.Dir(fileName)

	// 创建目录（如果目录不存在）
	if e := os.MkdirAll(dirName, 0755); e != nil {
		fmt.Fprintf(os.Stderr, "Failed to create directory: %s\n", err)
		return
	}

	// 尝试打开或创建文件
	file, e := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Failed to open/create log file: %s\n", err)
		return
	}

	// 定义一个函数来设置日志输出到控制台和文件
	setLogOutput := func(logger *Logger, output io.Writer) {
		logger.log.SetOutput(io.MultiWriter(file, stdio.Out)) // 同时输出到文件和控制台
	}

	// 根据日志级别设置日志输出
	switch level {
	case DEBUG:
		setLogOutput(dbg, file)
	case INFO:
		setLogOutput(info, file)
	case WARN:
		setLogOutput(warn, file)
	case ERROR:
		setLogOutput(err, file)
	case NONE:
		return
	default:
		return
	}
}
func Debug() *Logger {
	return dbg
}

func LogString(level Level, s string) string {
	var buffer bytes.Buffer
	// 将 logger 的输出改为 buffer
	l := log.New(&buffer, "", log.Ldate|log.Ltime|log.Llongfile)
	switch level {
	case DEBUG:
		l.SetPrefix("#DEBUG")
	case INFO:
		l.SetPrefix("#INFO")
	case WARN:
		l.SetPrefix("#WARN")
	case ERROR:
		l.SetPrefix("#ERROR")
	case DATA:
		l.SetPrefix("#DATA")
	}
	l.Println(s)
	return Clear(buffer.String())
}
