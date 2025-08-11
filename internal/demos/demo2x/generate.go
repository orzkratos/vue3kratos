package demo2x

import (
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
)

func GetProjectPath() string {
	return osmustexist.ROOT(runpath.PARENT.Join("vue3npm"))
}
