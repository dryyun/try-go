package initpkg4

import "try-go/ttools"

func init() {
	ttools.EchoFileLine("initpkg4 被 initpkg2 依赖，也被 init_try 依赖，只初始化一次")
	ttools.EchoFileLine("initpkg4 会在 initpkg2 之前初始化")
}
