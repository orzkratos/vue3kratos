package main

import (
	"os"
	"path/filepath"

	"github.com/orzkratos/vue3kratos"
	"github.com/orzkratos/vue3kratos/internal/demos/demo1x"
	"github.com/orzkratos/vue3kratos/internal/demos/demo2x"
	"github.com/orzkratos/vue3kratos/internal/utils"
	"github.com/yyle88/done"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
)

func main() {
	projectPath := demo1x.GetProjectPath()
	zaplog.SUG.Debugln(projectPath)

	outputRoot := filepath.Join(projectPath, "bin", "web_api_grpc_ts.out")
	zaplog.SUG.Debugln(outputRoot)
	done.Done(os.RemoveAll(outputRoot))

	done.VAE(osexec.ExecInPath(projectPath, "make", "web-api-grpc-ts")).Nice()
	osmustexist.MustRoot(outputRoot)

	vueRpcPath := filepath.Join(demo2x.GetProjectPath(), "src/rpc")
	osmustexist.MustRoot(vueRpcPath)
	utils.CopyFiles(outputRoot, vueRpcPath)

	vue3kratos.GenGrpcViaHttpInRoot(vueRpcPath)
}
