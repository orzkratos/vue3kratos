package vue3kratos2gen

import (
	"os"
	"strings"

	"github.com/orzkratos/vue3kratos/internal/utils"
	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func GenGrpcViaHttpInRoot(
	execInRoot string, //这个是执行代码转换的命令行所在的目录，当然在本项目里是 vue3npm 的目录的绝对路径
	grpcTsRoot string, //这个是存放以 ".client.ts" 结尾的代码文件的目录/代码总目录，注意不要填错，毕竟它会修改文件内容，这里要特别注意
) {
	zaplog.LOG.Info("gen_grpc_via_http", zap.String("exec_in_root", execInRoot))
	osmustexist.MustRoot(execInRoot)

	zaplog.LOG.Info("gen_grpc_via_http", zap.String("grpc_ts_root", grpcTsRoot))
	osmustexist.MustRoot(grpcTsRoot)

	done.Done(utils.WalkFiles(grpcTsRoot, func(path string, info os.FileInfo) error {
		zaplog.LOG.Info("walk", zap.String("path", path))
		if strings.HasSuffix(path, ".client.ts") {
			zaplog.LOG.Info("walk", zap.String("name", info.Name()))
			GenGrpcViaHttpInPath(execInRoot, path)
		}
		return nil
	}))
}

func GenGrpcViaHttpInPath(execInRoot string, grpcTsPath string) {
	// 这里执行的命令大概是这样的，在运行时也会在日志中打印出命令内容
	// cd xxx && npm run rpcrewrite -- /xxx/src/rpc/rpc_admin_login/admin_login.client.ts
	// 当然里面涉及的都是绝对路径
	// 这里不成功就直接崩掉，毕竟只是代码生成的逻辑，而不是在服务中运行的逻辑，直接崩掉也是可以的
	output := done.VAE(utils.ExecInPath(execInRoot, "npm", "run", "rpcrewrite", "--", grpcTsPath)).Nice()
	// 简单打印下运行的结果
	zaplog.SUG.Info(string(output))
}
