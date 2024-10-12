在这里使用 kratos 的 helloworld 项目，就可以体验工具。

https://go-kratos.dev/docs/getting-started/start/

根据官方网站的指导，在当前目录里，创建 helloworld 项目。

使用 make all 和 make build 把它跑起来。

接着修改其Makefile的内容。

我在这里放了个 Makefile，使用它，替换 helloworld 里面的（但是我们也知道 kratos 项目会升级，我的工具也会升级，因此这里的 [Makefile](Makefile) 只在写这个文档的时候是匹配的，以后只能当作参考而不能直接覆盖内容）。

接着使用
```
make web_api_grpc_ts
```
就可以使用这个文件里的子命令得到 typescript 的客户端代码。

这里为了便于展示效果，就把结果拷贝了出来

这是直接得到的 ts 代码 [转换前的客户端代码](greeter.client.ts-B.txt)

接着使用
```
vue3kratos2main gen-grpc-via-http-in-path --grpc_ts_path=/xxx/helloworld/bin/web_api_grpc_ts.out/helloworld/v1/greeter.client.ts
```
注意这里要用绝对路径

就会得到你想要的结果

这里为了便于展示效果，就把结果拷贝了出来

这是经过转换的 ts 代码 [转换后的客户端代码](greeter.client.ts-A.txt)

当然其他 ts 代码也同样重要，只是它们不会受到本工具的影响，因此就不展示啦。把整个生成的 `web_api_grpc_ts.out` 里的内容都拷贝到你的 vue 项目里，就可以使用啦。
