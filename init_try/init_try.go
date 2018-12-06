package main

import (
	"try-go/init_try/fileline"
	_ "try-go/init_try/initpkg1"
	_ "try-go/init_try/initpkg2"
	_ "try-go/init_try/initpkg3"
	_ "try-go/init_try/initpkg4"
)

func main() {
	fileline.EchoFl("最晚执行的 main ")
}
