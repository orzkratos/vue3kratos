package demo1x

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/orzkratos/vue3kratos"
	"github.com/orzkratos/vue3kratos/internal/utils"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath/osmustexist"
)

func TestGetProjectPath(t *testing.T) {
	t.Log(GetProjectPath())
}

func TestGetClientBPath(t *testing.T) {
	t.Log(GetClientBPath())
}

func TestGetClientAPath(t *testing.T) {
	t.Log(GetClientAPath())
}

// 测试在golang侧转换
func TestGenGrpcViaHttpInRoot(t *testing.T) {
	//转换前的代码
	dataB := done.VAE(os.ReadFile(GetClientBPath())).Nice()
	//转换后的代码
	dataA := done.VAE(os.ReadFile(GetClientAPath())).Nice()

	targetRoot, err := os.MkdirTemp("", utils.NewUUID())
	require.NoError(t, err)
	defer func() {
		done.Done(os.RemoveAll(targetRoot))
	}()
	osmustexist.ROOT(targetRoot)

	// 把转换前的代码拷贝到新文件里
	targetPath := filepath.Join(targetRoot, utils.NewUUID()+".client.ts")
	// 把转换前的文件内容写到文件里
	done.Done(os.WriteFile(targetPath, dataB, 0644))
	// 确保已经写入文件，即目标存在
	osmustexist.FILE(targetPath)

	t.Log("run generate")
	vue3kratos.GenGrpcViaHttpInRoot(targetRoot)
	t.Log("success done")

	//读取转换后的内容
	newContent := done.VAE(os.ReadFile(targetPath)).Nice()
	//确保和期望的相同
	require.Equal(t, string(dataA), string(newContent))
}
