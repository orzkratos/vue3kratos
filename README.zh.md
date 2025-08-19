# vue3kratos

> Vue 3 前端 + Kratos 后端集成开发工具包  
> 从 Go Kratos proto 文件生成类型安全的 TypeScript 客户端  
> 无缝的 TypeScript + Go RPC 集成，使用 `@protobuf-ts/plugin`  
> 从后端到前端的完整类型安全  
> 像调用本地函数一样调用后端 API

## 英文文档

[ENGLISH README](README.md)

## 核心架构

`vue3kratos` 连接 Go 后端与 Vue 3 前端，生成 TypeScript 客户端。

### 开发工具链
```
+-------------+    +----------+    +---------------+    +--------------+    +---------------+
| .proto 文件 | -> | protoc   | -> | gRPC TS 客户端| -> | vue3kratos   | -> | HTTP TS 客户端|
|             |    | + 插件   |    |               |    | CLI 转换工具 |    |               |
+-------------+    +----------+    +---------------+    +--------------+    +---------------+
```

### 运行时架构
```
+------------------+                                    +------------------+
|   Vue 3 前端     |                                    |  Kratos 后端     |
|                  |                                    |                  |
| ┌──────────────┐ |                                    | ┌──────────────┐ |
| │ gRPC客户端代码│ |  代码编写 gRPC 风格调用             │ │   HTTP 服务   │ |
| │ client.call() │ |                                    | │  :28000       │ |
| └──────┬───────┘ |                                    | └──────────────┘ |
|        │         |                                    |                  |
| ┌──────▼───────┐ |   底层实际发送 HTTP 请求            |                  |
| │HTTP 转换层    │ |  =========================>        |                  |
| │(@yyle88/grpt) │ |     POST /api/method               |                  |
| └──────────────┘ |     Content-Type: application/json |                  |
|                  |                                    | ┌──────────────┐ |
|                  |                                    | │   gRPC 服务   │ |
|                  |                    ❌ 跳过 ------> | │   :28001      │ |
|                  |                                    | │  (未使用)     │ |
|                  |                                    | └──────────────┘ |
+------------------+                                    +------------------+
```

## 🌟 核心特性

*   **自动代码生成**: 从 proto 文件生成简洁的 TypeScript 客户端
*   **告别手写API**: 不再需要手动编写API客户端代码
*   **一键转换**: 单条命令将 gRPC 客户端转为 HTTP 客户端
*   **Web 兼容**: 直接在 Web 中使用，无需 gRPC 复杂性
*   **完整类型安全**: 从后端到前端的完整类型检查
*   **IDE 智能提示**: 丰富的开发体验，智能代码补全
*   **Makefile 集成**: 轻松集成到现有构建流程
*   **CI/CD 管线**: 平滑的工作流自动化支持
*   **Axios HTTP 客户端**: 现代化的 HTTP 客户端实现
*   **本地函数体验**: 像调用本地函数一样调用 API

## 🚀 快速上手

在几分钟内，亲身体验 `vue3kratos` 的强大之处。

### 1. 启动后端服务

首先，运行 `demo1go` 后端服务。

```bash
# 进入后端服务DIR
cd internal/demos/demo1x/demo1go

# 启动服务
make run
# 或者
go run ./cmd/demo1go
```

服务将在 `http://127.0.0.1:28000` (HTTP) 和 `http://127.0.0.1:28001` (gRPC) 上启动。

### 2. 运行前端演示

接着，在另一个终端中，运行 `vue3npm` 前端演示。

```bash
# 进入前端演示DIR
cd internal/demos/demo2x/vue3npm

# 安装依赖
npm install

# 运行智慧演示
npm run demo:wise
```

现在，你将看到前端应用调用后端API的实时日志输出，完美地展示了前后端的无缝协作。

## `vue3kratos` 工作流详解

在你的项目中使用 `vue3kratos` 只需以下几步。

### 第1步: 安装链

确保你已全局安装 `@protobuf-ts/plugin`，并已安装 `vue3kratos` 的Go CLI应用。

```bash
# 安装 protobuf-ts 插件（使用这个，而不是其他 protoc-gen-ts 插件）
npm install -g @protobuf-ts/plugin

# 确认是否安装成功
which protoc-gen-ts

# 安装 vue3kratos CLI
go install github.com/orzkratos/vue3kratos/cmd/vue3orzkratos@latest
```

#### 基础开发环境示例

```bash
# 示例全局依赖
npm list -g --depth=0
├── @protobuf-ts/plugin@2.9.4
├── typescript@5.4.5
├── vite@5.4.8
```

### 第2步: 生成 gRPC 客户端

在你的 Kratos 项目中，使用 `protoc` 和 `protoc-gen-ts` 从 `.proto` 文件生成 TypeScript 客户端。推荐在 `Makefile` 中进行管理。

```makefile
web_api_grpc_ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	PROTOC_GEN_TS=$$(which protoc-gen-ts) && \
	protoc \
	--plugin=protoc-gen-ts=$$PROTOC_GEN_TS \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./api \
	--proto_path=./third_party \
	$(API_PROTO_FILES)

	PROTOC_GEN_TS=$$(which protoc-gen-ts) && \
	protoc \
	--plugin=protoc-gen-ts=$$PROTOC_GEN_TS \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./third_party \
	$(THIRD_PARTY_GOOGLE_API_PROTO_FILES)
```

Makefile 中还需添加：

```makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
```

### 第3步: 转换为 HTTP 客户端

使用 `vue3orzkratos` CLI 将上一步生成的 gRPC 客户端文件转换为 HTTP 客户端。

```bash
vue3orzkratos gen-grpc-via-http-in-path --grpc-ts-path=/path/to/the/generated.client.ts
```

该命令会将目标文件中 gRPC 的调用方式替换为基于 Axios 的 HTTP 请求。

### 第4步: 在 Vue 中使用

在你的 Vue 项目中，安装 `@yyle88/grpt` 辅助模块，然后你就可以像调用一个普通函数一样，调用所有API。

```bash
npm install @yyle88/grpt
```

> npm 模块地址：[@yyle88/grpt](https://www.npmjs.com/package/@yyle88/grpt)

```typescript
// 来自项目实际示例 internal/demos/demo2x/vue3npm/src/demo-ping.ts
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { RpcpingClient } from "./rpc/rpcping/rpcping.client";
import { StringValue } from "./rpc/google/protobuf/wrappers";

// 创建传输实例
const demoTransport = new GrpcWebFetchTransport({
    baseUrl: "http://127.0.0.1:28000",
    meta: {
        Authorization: 'TOKEN-888',
    },
});

const rpcpingClient = new RpcpingClient(demoTransport);

// API 调用示例
async function demoPing() {
    const request = StringValue.create({
        value: "Hello from Vue3 Kratos!"
    });

    try {
        const response = await rpcpingClient.ping(request, {});
        console.log('Ping 成功:', response.data.value);
        return response.data.value;
    } catch (err) {
        console.error('Ping 失败:', err);
        throw err;
    }
}
```

---

## 🔁 示例项目

* **[internal/demos/demo1x/demo1go](internal/demos/demo1x/demo1go)** – Kratos 项目生成 TypeScript 客户端代码
* **[internal/demos/demo2x/vue3npm](internal/demos/demo2x/vue3npm)** – HTTP 转换和 Vue 3 集成示例

---

## ✅ 特性小结

* 将 Kratos proto 文件生成为类型安全的 TypeScript gRPC 客户端
* 支持自动替换为 HTTP 请求形式（Axios 封装）
* 支持类型提示与自动补全，前端集成体验极佳
* 可集成至 Makefile 或 CI/CD 中
* Web 兼容的 HTTP 客户端，支持直接前端调用

---

## 💡 适合人群

* 使用 Kratos 作为后端的开发者
* 想通过 Vue 3 编写前端，直接调用后端服务
* 希望实现“全类型安全”的调用体验

---

## 旧版本文档

[旧版说明](internal/doc/README_OLD_DOC.zh.md)

---

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
