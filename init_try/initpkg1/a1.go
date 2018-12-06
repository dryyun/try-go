package initpkg1

import (
	"try-go/init_try/fileline"
)

var initpkg1_a1_0 int = pkg1a10()
var initpkg1_a1_a int = initpkg1_a1_b + pkg1a1a()
var initpkg1_a1_b int = pkg1a1b()

func init() {
	// init 1
	fileline.EchoFl("同包变量初始化完成，现在是 init ，多.go 文件，按照顺序先初始化 a1.go")
	fileline.EchoFl("同文件多 init，按照出现顺序初始化 a.go - a1.init1")
}

func init() {
	fileline.EchoFl("同文件多 init，按照出现顺序初始化 a.go - a1.init2")
}

func pkg1a10() int {
	fileline.EchoFl("初始化第一个包 initpkg1 的第一个顺序文件 a1.go")
	fileline.EchoFl("同包中先初始化变量、常量 ，然后才是 init 函数，a1.initpkg1_a1_0 ")
	return 0
}

func pkg1a1a() int {
	fileline.EchoFl("变量依赖，a1.initpkg1_a1_a，之前 a1.initpkg1_a1_b 已经被初始化了")
	return 0
}

func pkg1a1b() int {
	fileline.EchoFl("初始化变量依赖，先初始化 a1.initpkg1_a1_b ，然后才是 a1.initpkg1_a1_a ")
	return 0
}
