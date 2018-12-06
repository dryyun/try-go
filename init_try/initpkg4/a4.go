package initpkg4

import "try-go/init_try/fileline"

func init() {
	fileline.EchoFl("initpkg4 被 initpkg2 依赖，也被 init_try 依赖，只初始化一次")
	fileline.EchoFl("initpkg4 会在 initpkg2 之前初始化")
}
