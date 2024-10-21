package vue3kratos2gen

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
)

func TestGenGrpcViaHttpInCode(t *testing.T) {
	srcContent := done.VAE(os.ReadFile(osmustexist.FILE(runpath.PARENT.Join("./../internal/demos/greeter.client.ts-B.txt")))).Nice()
	t.Log(string(srcContent))

	newContent := done.VAE(os.ReadFile(osmustexist.FILE(runpath.PARENT.Join("./../internal/demos/greeter.client.ts-A.txt")))).Nice()
	t.Log(string(newContent))

	// 比较目标文件内容和生成的内容是否一致，使用 string 比较内容，当有不同时能够很方便看到
	require.Equal(t, string(newContent), GenGrpcViaHttpInCode(string(srcContent)))
}
