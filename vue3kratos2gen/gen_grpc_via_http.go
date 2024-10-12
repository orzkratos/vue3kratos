package vue3kratos2gen

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/orzkratos/vue3kratos/internal/utils"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// GenGrpcViaHttpInRoot 在整个目录里找 ts grpc client 代码，把它们转换为使用 http 请求
// 这个是存放以 ".client.ts" 结尾的代码文件的目录/代码总目录，注意不要填错，毕竟它会修改文件内容，这里要特别注意
func (G *GenType) GenGrpcViaHttpInRoot(grpcTsRoot string) {
	zaplog.LOG.Info("gen_grpc_via_http", zap.String("exec_in_root", G.execInRoot))
	osmustexist.MustRoot(G.execInRoot)

	zaplog.LOG.Info("gen_grpc_via_http", zap.String("grpc_ts_root", grpcTsRoot))
	osmustexist.MustRoot(grpcTsRoot)

	done.Done(utils.WalkFiles(grpcTsRoot, func(path string, info os.FileInfo) error {
		zaplog.LOG.Info("walk", zap.String("path", path))
		if strings.HasSuffix(path, ".client.ts") {
			zaplog.LOG.Info("walk", zap.String("name", info.Name()))

			switch G.wkLanguage {
			case GO:
				done.Done(GenGrpcViaHttpWithGo(path))
			case JS:
				GenGrpcViaHttpUseNpm(G.execInRoot, path)
			default:
				zaplog.LOG.Panic("wrong param")
			}
		}
		return nil
	}))
}

func GenGrpcViaHttpUseNpm(execInRoot string, codePath string) {
	// 确保工作目录是项目目录
	osmustexist.File(filepath.Join(execInRoot, "package.json"))
	// 确保文件路径也是存在的
	osmustexist.File(codePath)
	// 这里执行的命令大概是这样的，在运行时也会在日志中打印出命令内容
	// cd xxx && npm run rpcrewrite -- /xxx/src/rpc/rpc_admin_login/admin_login.client.ts
	// 当然里面涉及的都是绝对路径
	// 这里不成功就直接崩掉，毕竟只是代码生成的逻辑，而不是在服务中运行的逻辑，直接崩掉也是可以的
	output := done.VAE(utils.ExecInPath(execInRoot, "npm", "run", "rpcrewrite", "--", codePath)).Nice()
	// 简单打印下运行的结果
	zaplog.SUG.Info(string(output))
}

func GenGrpcViaHttpWithGo(codePath string) error {
	zaplog.LOG.Info("gen_grpc_via_http", zap.String("code_path", codePath))
	if codePath == "" {
		return erero.New("没有路径参数-请设置文件路径.")
	}
	// 读取文件
	newContent := string(done.VAE(os.ReadFile(codePath)).Nice())

	// 进行替换，把逻辑里调用grpc的地方改为调用http
	newContent = strings.Replace(newContent, "stackIntercept<", "executeGrtp<", -1)
	newContent = strings.Replace(newContent, "UnaryCall<", "GrtpPromise<", -1)

	// 判断是否已经存在目标引用
	targetImport := `import { executeGrtp, GrtpPromise } from '@/rpcviahttp/rpcviahttp.ts';`
	searchImport := `import type { RpcOptions } from "@protobuf-ts/runtime-rpc";`

	if !strings.Contains(newContent, targetImport) {
		// 找到指定引用的位置
		targetIndex := strings.Index(newContent, searchImport)
		if targetIndex != -1 {
			// 在目标引用的后面添加新的引用
			insertIndex := targetIndex + len(searchImport)
			// 拼接出结果
			newContent = newContent[:insertIndex] + "\n" + targetImport + newContent[insertIndex:]
		}
	}

	// 写回文件
	done.Done(os.WriteFile(codePath, []byte(newContent), 0644))
	zaplog.LOG.Info("内容替换成功!")
	return nil
}
