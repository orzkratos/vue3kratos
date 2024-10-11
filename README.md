# vue3kratos
Vue 3 frontends and Kratos backends integration toolkit. Enabling seamless communication and efficient full-stack development.

做后端做久了偶尔也要做做前端。

在学前端时发现vue3用起来还行，就想着用vue3做前端连接kratos的后端服务，因此做了个中间的胶水工具，让两个语言能够更顺畅的对接。

## 把kratos的proto接口定义转换为 typescript 语言的 grpc 的请求
```
web_api_grpc_ts:
        mkdir -p ./bin/web_api_grpc_ts.out
        protoc \
      --plugin=protoc-gen-ts=/home/yangyile/.nvm/versions/node/v20.14.0/bin/protoc-gen-ts \
      --ts_out=./bin/web_api_grpc_ts.out \
      --proto_path=./api \
      --proto_path=./proto3ps \
      $(API_PROTO_FILES)

        protoc \
      --plugin=protoc-gen-ts=/home/yangyile/.nvm/versions/node/v20.14.0/bin/protoc-gen-ts \
      --ts_out=./bin/web_api_grpc_ts.out \
      --proto_path=./proto3ps \
      $(THIRD_PARTY_GOOGLE_API_PROTO_FILES)
```

## 把grpc请求逻辑替换为使用http请求逻辑
安装
```
go install github.com/orzkratos/vue3kratos/cmd/vue3kratos2main@latest
```

使用
```
vue3kratos2main gen-grpc-via-http-in-path --grpc_ts_path=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
```
这个文件里的内容就会被替换，因此特别注意，路径要传正确的绝对路径。

结果里要引用的模块在这里: [把typescript的grpc请求转为http请求](vue3npm/src/rpcviahttp/rpcviahttp.ts)
