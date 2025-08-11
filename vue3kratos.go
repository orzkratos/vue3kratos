package vue3kratos

import (
	"os"
	"strings"

	"github.com/orzkratos/vue3kratos/internal/utils"
	"github.com/yyle88/done"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// GenGrpcViaHttpInRoot 在整个目录里找 ts grpc client 代码，把它们转换为使用 http 请求
// 这个是存放以 ".client.ts" 结尾的代码文件的目录/代码总目录，注意不要填错，毕竟它会修改文件内容，这里要特别注意
func GenGrpcViaHttpInRoot(grpcTsRoot string) {
	zaplog.LOG.Info("gen-grpc-via-http", zap.String("grpc-ts-root", grpcTsRoot))
	osmustexist.MustRoot(grpcTsRoot)

	done.Done(utils.WalkFiles(grpcTsRoot, func(path string, info os.FileInfo) error {
		zaplog.LOG.Info("walk", zap.String("path", path))
		if strings.HasSuffix(path, ".client.ts") {
			zaplog.LOG.Info("walk", zap.String("name", info.Name()))
			GenGrpcViaHttpInPath(path)
		}
		return nil
	}))
}

// GenGrpcViaHttpInPath 在整个文件里找 ts grpc client 代码，把它们转换为使用 http 请求
// 把它替换完将会直接写进原来的文件里
func GenGrpcViaHttpInPath(grpcTsPath string) {
	zaplog.LOG.Info("gen-grpc-via-http", zap.String("code_path", grpcTsPath))
	must.Nice(grpcTsPath)
	// 读取文件
	srcContent := string(done.VAE(os.ReadFile(grpcTsPath)).Nice())

	// 替换内容-得到新的内容
	newContent := GenGrpcViaHttpInCode(srcContent)

	// 写回文件
	done.Done(os.WriteFile(grpcTsPath, []byte(newContent), 0644))
	zaplog.LOG.Info("content replace success!!!")
}

// GenGrpcViaHttpInCode 在整个代码里找 ts grpc client 代码，把它们转换为使用 http 请求
// 将返回修改后的代码
// 假如文件已经被替换，再次运行时就不改变文件内容，保持替换后的状态
func GenGrpcViaHttpInCode(srcContent string) string {
	newContent := srcContent
	// 进行替换，把逻辑里调用grpc的地方改为调用http
	newContent = strings.ReplaceAll(newContent, "stackIntercept<", "executeGrtp<")
	newContent = strings.ReplaceAll(newContent, "UnaryCall<", "GrtpPromise<")

	// 判断是否已经存在目标引用
	targetImport := `import { executeGrtp } from '@yyle88/grpt/src/grpcviahttp';` +
		"\n" +
		`import type { GrtpPromise } from '@yyle88/grpt/src/grpcviahttp';`
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
	return newContent
}

// CloneFilesToDestRoot recursively copies files from source to target.
// 把源目录的文件克隆到目标目录里
func CloneFilesToDestRoot(sourceRoot string, targetRoot string) {
	utils.CopyFiles(sourceRoot, targetRoot)
}
