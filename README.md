# vue3kratos
Vue 3 frontends and Kratos backends integration toolkit. Enabling seamless communication and efficient full-stack development.

做后端做久了偶尔也要做做前端。

在学前端时发现vue3用起来还行，就想着用vue3做前端连接kratos的后端服务。

因此做了个中间的胶水工具，让两个语言能够更顺畅的对接。

## 把kratos的proto接口定义转换为 typescript 语言的 grpc 的 client 客户端代码
### 使用的工具链安装
主要是这里：
[需要使用的 protoc-gen-ts 工具网页](https://www.npmjs.com/package/protoc-gen-ts)
很明显的，里面有 usage 的指导：
```
npm install -g protoc-gen-ts
```

### 工具链安装的操作
```
#需要看看是否已安装：
admin@lele-de-MacBook-Pro ~ % which protoc-gen-ts
protoc-gen-ts not found

#假如没有安装就安装：
admin@lele-de-MacBook-Pro ~ % npm install -g protoc-gen-ts

added 1 package in 4s

1 package is looking for funding
  run `npm fund` for details

#安装完毕以后看位置：
admin@lele-de-MacBook-Pro ~ % which protoc-gen-ts         
/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts

#安装的位置非常重要。
```

### 使用工具生成代码
#### 根据 proto 得到 typescript grpc 的 client 客户端代码
注意：这里只能给 makefile 的参考内容，因为开发者的工具路径不同，三方包的路径和版本不同

#### 配置样例1:
这是我的某个环境的逻辑，需要你在 kratos 项目的 Makefile 里面增加这段：
``` makefile
web_api_grpc_ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	protoc \
	--plugin=protoc-gen-ts=/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./api \
	--proto_path=./third_party \
	$(API_PROTO_FILES)

	protoc \
	--plugin=protoc-gen-ts=/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./third_party \
	$(THIRD_PARTY_GOOGLE_API_PROTO_FILES)
```
你需要在`Makefile`找路径的逻辑里增加(这还只是mac/ubuntu系统的，windows系统的你自己也写写吧):
``` makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
```

#### 配置样例2:
这是我的其它环境的逻辑，需要你在 kratos 项目的 Makefile 里面增加这段：
``` makefile
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
你需要在`Makefile`找路径的逻辑里增加(这还只是mac/ubuntu系统的，windows系统的你自己也写写吧):
``` makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find proto3ps/google/api -name *.proto)
```

具体使用时请自己根据实际编写吧，确保工具链相同。

#### 通过工具得到代码
```
make web_api_grpc_ts
```
假如不报错的话就会得到你想要的结果代码
```
admin@lele-de-MacBook-Pro helloworld % make web_api_grpc_ts
admin@lele-de-MacBook-Pro helloworld % 
admin@lele-de-MacBook-Pro helloworld % cd bin/web_api_grpc_ts.out 
admin@lele-de-MacBook-Pro web_api_grpc_ts.out % ls
google          helloworld
admin@lele-de-MacBook-Pro web_api_grpc_ts.out %
```
但是很明显的，这个代码是基于grpc协议的，假如你需要把代码改为底层使用http协议传输，也有多种解决方案。

比如使用 web-grpc 代理等，但是需要额外的配置。

其中我认为最简单的，就是把客户端底层代码改改，让调用grpc的变为调用http接口，在底层使用`axios`请求。

这样的好处是，以函数调用的形式请求http接口，而且参数和返回值都是 typescript 带类型的，就非常的方便。

因此接下来就是使用这个工具把代码再改改。

## 把grpc请求逻辑替换为使用http请求逻辑
### 安装
```
go install github.com/orzkratos/vue3kratos/cmd/vue3kratos2main@latest
```

### 使用
```
vue3kratos2main gen-grpc-via-http-in-path --grpc_ts_path=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
```
这个文件里的内容就会被替换，因此特别注意，路径要传正确的绝对路径。

结果里要引用的模块在这里: [把typescript的grpc请求转为http请求](vue3npm/src/rpcviahttp/rpcviahttp.ts)
