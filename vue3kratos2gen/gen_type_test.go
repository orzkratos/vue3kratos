package vue3kratos2gen

import (
	"testing"

	"github.com/yyle88/runpath"
)

func TestNewGen(t *testing.T) {
	gen := NewGen().SetExecInRoot(runpath.PARENT.Path())
	t.Log(gen.execInRoot)
	t.Log(gen.wkLanguage)
}
