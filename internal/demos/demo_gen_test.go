package demos

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/orzkratos/vue3kratos/internal/utils"
	"github.com/orzkratos/vue3kratos/vue3kratos2gen"
	"github.com/orzkratos/vue3kratos/vue3npm"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
)

// 测试在golang侧转换
func TestGenWithGo(t *testing.T) {
	caseGen(t, func(targetRoot string, targetPath string) {
		vue3kratos2gen.GenGrpcViaHttpInRoot(targetRoot)
	})
}

// 测试在npm js侧转换，这里用的命令行，调用npm项目里的命令做转换
func TestGenUseNpm(t *testing.T) {
	npmRunRoot := osmustexist.ROOT(vue3npm.SourceRoot())
	// 确保工作目录是项目目录
	osmustexist.FILE(filepath.Join(npmRunRoot, "package.json"))

	caseGen(t, func(targetRoot string, targetPath string) {
		// 这里执行的命令大概是这样的，在运行时也会在日志中打印出命令内容
		// cd xxx && npm run rpcrewrite -- /xxx/src/rpc/rpc_admin_login/admin_login.client.ts
		// 当然里面涉及的都是绝对路径
		// 这里不成功就直接崩掉，毕竟只是代码生成的逻辑，而不是在服务中运行的逻辑，直接崩掉也是可以的
		output := done.VAE(utils.ExecInPath(npmRunRoot, "npm", "run", "rpcrewrite", "--", targetPath)).Nice()
		// 简单打印下运行的结果
		zaplog.SUG.Info(string(output))
	})
}

func caseGen(t *testing.T, runGen func(targetRoot string, targetPath string)) {
	//转换前的代码
	dataB := done.VAE(os.ReadFile(runpath.PARENT.Join("greeter.client.ts-B.txt"))).Nice()
	//转换后的代码
	dataA := done.VAE(os.ReadFile(runpath.PARENT.Join("greeter.client.ts-A.txt"))).Nice()

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

	t.Log("gen-run")
	runGen(targetRoot, targetPath)
	t.Log("success")

	//读取转换后的内容
	newContent := done.VAE(os.ReadFile(targetPath)).Nice()
	//确保和期望的相同
	require.Equal(t, string(dataA), string(newContent))
}
