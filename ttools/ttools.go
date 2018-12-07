package ttools

import (
	"fmt"
	"runtime"
)

func EchoFileLine(s string) {
	_, filename, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("文件= %s，行号= %d，\t输出= %s\n\n", filename, line, s)
	}
}
