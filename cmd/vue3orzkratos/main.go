package main

import (
	"fmt"
	"os"

	"github.com/orzkratos/vue3kratos"
	"github.com/spf13/cobra"
	"github.com/yyle88/done"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "vue3orzkratos",
		Short: "A CLI app to the vue3kratos chain.",
		Long:  `vue3orzkratos is a command-line app that helps bridge the gap between Kratos gRPC backends and Vue 3 frontends by converting generated TypeScript clients to be web-safe.`,
	}

	// example: vue3orzkratos gen-grpc-via-http-in-root --grpc-ts-root=/xxx/src/rpc
	rootCmd.AddCommand(makeRootCmd())

	// example: vue3orzkratos gen-grpc-via-http-in-path --grpc-ts-path=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
	rootCmd.AddCommand(makePathCmd())

	// example: vue3orzkratos gen-grpc-via-http-in-code --grpc-ts-code=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
	rootCmd.AddCommand(makeCodeCmd())

	// Execute root command
	done.Done(rootCmd.Execute())
}

func makeRootCmd() *cobra.Command {
	var grpcTsRoot string
	var genCmd = &cobra.Command{
		Use:   "gen-grpc-via-http-in-root",
		Short: "Deep finds and converts each .client.ts file in a DIR.",
		Long:  `Searches to files ending with .client.ts within the specified root DIR and all its subdirectories, then converts each file from a gRPC client to an HTTP client.`,
		Run: func(cmd *cobra.Command, args []string) {
			vue3kratos.GenGrpcViaHttpInRoot(grpcTsRoot)
			zaplog.LOG.Info("✅ Success! Each found .client.ts file has been converted.", zap.String("root", grpcTsRoot))
		},
	}
	genCmd.Flags().StringVarP(&grpcTsRoot, "grpc-ts-root", "r", "", "Root DIR to search to .client.ts files.")
	done.Done(genCmd.MarkFlagRequired("grpc-ts-root"))
	return genCmd
}

func makePathCmd() *cobra.Command {
	var grpcTsPath string
	var genCmd = &cobra.Command{
		Use:   "gen-grpc-via-http-in-path",
		Short: "Converts one .client.ts file.",
		Long:  `Converts one specified .client.ts file from a gRPC client to an HTTP client, modifying it in place.`,
		Run: func(cmd *cobra.Command, args []string) {
			vue3kratos.GenGrpcViaHttpInPath(grpcTsPath)
			zaplog.LOG.Info("✅ Success! File has been converted.", zap.String("path", grpcTsPath))
		},
	}
	genCmd.Flags().StringVarP(&grpcTsPath, "grpc-ts-path", "p", "", "Path to the target .client.ts file.")
	done.Done(genCmd.MarkFlagRequired("grpc-ts-path"))
	return genCmd
}

func makeCodeCmd() *cobra.Command {
	var grpcTsCodePath string
	var genCmd = &cobra.Command{
		Use:   "gen-grpc-via-http-in-code",
		Short: "Converts and prints the content of a .client.ts file to stdout.",
		Long:  `Reads a .client.ts file, performs the conversion in memory, and prints the resulting code to standard output without modifying the source file.`,
		Run: func(cmd *cobra.Command, args []string) {
			clientPath := osmustexist.FILE(grpcTsCodePath)
			srcContent := done.VAE(os.ReadFile(clientPath)).Nice()
			newContent := vue3kratos.GenGrpcViaHttpInCode(string(srcContent))
			fmt.Println(newContent)
		},
	}
	genCmd.Flags().StringVarP(&grpcTsCodePath, "grpc-ts-code", "c", "", "Path to the .client.ts file to read and convert.")
	done.Done(genCmd.MarkFlagRequired("grpc-ts-code"))
	return genCmd
}
