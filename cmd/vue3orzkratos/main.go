package main

import (
	"fmt"
	"os"

	"github.com/orzkratos/vue3kratos"
	"github.com/spf13/cobra"
	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
)

// 定义根命令
var rootCmd = &cobra.Command{
	Use:   "run", // 根命令的名称
	Short: "CLI: vue3kratos",
	Long:  "see: https://github.com/orzkratos/vue3kratos",
	Run: func(cmd *cobra.Command, args []string) {
		zaplog.LOG.Info("run")
	},
}

func main() {
	{ // example: vue3orzkratos gen-grpc-via-http-in-root --grpc_ts_root=/xxx/src/rpc
		// 目标的位置
		var grpcTsRoot string

		// 创建子命令
		var genCmd = &cobra.Command{
			Use:   "gen-grpc-via-http-in-root",
			Short: "Generate gRPC via HTTP",
			Long:  "Generate gRPC TypeScript stubs using HTTP",
			Run: func(cmd *cobra.Command, args []string) {
				vue3kratos.GenGrpcViaHttpInRoot(grpcTsRoot)
			},
		}
		// 把 grpcTsRoot 路径参数加入子命令里
		genCmd.Flags().StringVarP(&grpcTsRoot, "grpc_ts_root", "", "", "where is .client.ts in")

		// 设置为必须参数
		done.Done(genCmd.MarkFlagRequired("grpc_ts_root"))

		// 将子命令添加到根命令里
		rootCmd.AddCommand(genCmd)
	}

	{ // example: vue3orzkratos gen-grpc-via-http-in-path --grpc_ts_path=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
		// 目标的位置
		var grpcTsPath string

		// 创建子命令
		var genCmd = &cobra.Command{
			Use:   "gen-grpc-via-http-in-path",
			Short: "Generate gRPC via HTTP",
			Long:  "Generate gRPC TypeScript stubs using HTTP",
			Run: func(cmd *cobra.Command, args []string) {
				done.Done(vue3kratos.GenGrpcViaHttpInPath(grpcTsPath))
			},
		}
		// 把 grpcTsPath 路径参数加入子命令里
		genCmd.Flags().StringVarP(&grpcTsPath, "grpc_ts_path", "", "", "where is .client.ts in")

		// 设置为必须参数
		done.Done(genCmd.MarkFlagRequired("grpc_ts_path"))

		// 将子命令添加到根命令里
		rootCmd.AddCommand(genCmd)
	}

	{ // example: vue3orzkratos gen-grpc-via-http-in-code --grpc_ts_code=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
		// 目标的位置
		var grpcTsCode string

		// 创建子命令
		var genCmd = &cobra.Command{
			Use:   "gen-grpc-via-http-in-code",
			Short: "Generate gRPC via HTTP",
			Long:  "Generate gRPC TypeScript stubs using HTTP",
			Run: func(cmd *cobra.Command, args []string) {
				//这里需要的就是个文件的路径
				clientPath := osmustexist.FILE(grpcTsCode)
				//读取内容
				srcContent := done.VAE(os.ReadFile(clientPath)).Nice()
				//转换内容
				newContent := vue3kratos.GenGrpcViaHttpInCode(string(srcContent))
				//打印结果
				fmt.Println()
				fmt.Println(newContent)
				fmt.Println()
			},
		}
		// 把 grpcTsCode 路径参数加入子命令里
		genCmd.Flags().StringVarP(&grpcTsCode, "grpc_ts_code", "", "", "where is .client.ts in")

		// 设置为必须参数
		done.Done(genCmd.MarkFlagRequired("grpc_ts_code"))

		// 将子命令添加到根命令里
		rootCmd.AddCommand(genCmd)
	}

	// 执行根命令
	done.Done(rootCmd.Execute())
}
