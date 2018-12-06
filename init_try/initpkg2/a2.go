package initpkg2

import "try-go/init_try/fileline"
import _ "try-go/init_try/initpkg4"

func init() {
	fileline.EchoFl("不同包，按照字母顺序，初始化，initpkg2")
}
