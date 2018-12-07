package initpkg2

import (
	_ "try-go/init_try/initpkg4"
	"try-go/ttools"
)

func init() {
	ttools.EchoFileLine("不同包，按照字母顺序，初始化，initpkg2")
}
