package initpkg1

import "try-go/init_try/fileline"

var initpkg1_b1_a int = pkg1b1a()

func pkg1b1a() int {
	fileline.EchoFl("初始化第一个包 initpkg1 的第二个顺序文件 b1.go")
	fileline.EchoFl("同包中先初始化变量、常量，然后才是 init 函数，b1.initpkg1_b1_a")
	return 0
}

func init()  {
	fileline.EchoFl("同包中最晚初始化的 b1.init1 ，顺序 a1.init1 -> a1.init2 -> b1.init1 ")
}
