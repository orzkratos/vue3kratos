package vue3kratos2gen

import "github.com/yyle88/runpath"

type Language string

const (
	GO Language = "GO"
	JS Language = "JS"
)

type GenType struct {
	execInRoot string //这个是执行代码转换的命令行所在的目录，当然在本项目里是 vue3npm 的目录的绝对路径
	wkLanguage Language
}

func NewGen() *GenType {
	return &GenType{
		execInRoot: runpath.PARENT.Path(),
		wkLanguage: GO,
	}
}

func (G *GenType) SetExecInRoot(execInRoot string) *GenType {
	G.execInRoot = execInRoot
	return G
}

func (G *GenType) SetWkLanguage(wkLanguage Language) *GenType {
	G.wkLanguage = wkLanguage
	return G
}
