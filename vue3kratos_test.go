package vue3kratos

import (
	"os"
	"testing"

	"github.com/orzkratos/vue3kratos/internal/demos/demo1x"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath/osmustexist"
)

func TestGenGrpcViaHttpInCode(t *testing.T) {
	srcContent := done.VAE(os.ReadFile(osmustexist.FILE(demo1x.GetClientBPath()))).Nice()
	t.Log(string(srcContent))

	newContent := done.VAE(os.ReadFile(osmustexist.FILE(demo1x.GetClientAPath()))).Nice()
	t.Log(string(newContent))

	// 比较目标文件内容和生成的内容是否一致，使用 string 比较内容，当有不同时能够很方便看到
	require.Equal(t, string(newContent), GenGrpcViaHttpInCode(string(srcContent)))
}
