package main

import (
	"demo1go/api/rpcdemo"
	"demo1go/api/rpcping"
	"demo1go/api/v1/greeter"
	"demo1go/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/zapkratos"
	"github.com/yyle88/must"
	"github.com/yyle88/zaplog"
)

func main() {
	zapKratos := zapkratos.NewZapKratos(zaplog.LOGGER, zapkratos.NewOptions())
	slog := zapKratos.SubZap()
	slog.LOG.Debug("starting demo-service-kratos")

	// cp from https://github.com/go-kratos/examples/blob/61daed1ec4d5a94d689bc8fab9bc960c6af73ead/helloworld/server/main.go#L42
	httpSrv := http.NewServer(
		http.Address(":28000"),
		http.Middleware(
			recovery.Recovery(),
			logging.Server(zapKratos.GetLogger("HTTP")),
		),
	)

	// cp from https://github.com/go-kratos/examples/blob/61daed1ec4d5a94d689bc8fab9bc960c6af73ead/helloworld/server/main.go#L48
	grpcSrv := grpc.NewServer(
		grpc.Address(":28001"),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(zapKratos.GetLogger("GRPC")),
		),
	)

	pingService := service.NewRpcpingService(zapKratos.GetLogger("PING"))
	rpcping.RegisterRpcpingHTTPServer(httpSrv, pingService)
	rpcping.RegisterRpcpingServer(grpcSrv, pingService)

	demoService := service.NewRpcdemoService(zapKratos.GetLogger("DEMO"))
	rpcdemo.RegisterRpcdemoHTTPServer(httpSrv, demoService)
	rpcdemo.RegisterRpcdemoServer(grpcSrv, demoService)

	greeterService := service.NewGreeterService()
	greeter.RegisterGreeterHTTPServer(httpSrv, greeterService)
	greeter.RegisterGreeterServer(grpcSrv, greeterService)

	app := kratos.New(
		kratos.Name("demo-service-kratos"),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	slog.LOG.Debug("starting demo-service-kratos")
	must.Done(app.Run())
}
