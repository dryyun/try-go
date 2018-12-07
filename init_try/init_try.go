package main

import (
	_ "try-go/init_try/initpkg1"
	_ "try-go/init_try/initpkg2"
	_ "try-go/init_try/initpkg3"
	_ "try-go/init_try/initpkg4"
	"try-go/ttools"
)

func main() {
	ttools.EchoFileLine("最晚执行的 main ")
}
