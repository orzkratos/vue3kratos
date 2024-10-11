package vue3npm

import "github.com/yyle88/runpath"

// SourceRoot 获取到这个子项目的路径，以便于使用 "os/exec" 在这个目录下执行命令行操作
func SourceRoot() string {
	return runpath.PARENT.Path()
}
