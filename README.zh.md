# vue3kratos

> Vue 3 frontends + Kratos backends 集成开发工具包  
> Seamless TypeScript + Go RPC integration, powered by `@protobuf-ts/plugin` and Kratos

---

## 英文文档

[ENGLISH README](README.md)

---

## ✨ 项目简介

`vue3kratos` 是一个用于将 [Kratos](https://go-kratos.dev/) 后端服务定义（proto 文件）自动生成为 Vue 3 / TypeScript 客户端调用代码的工具链。  
它还支持将生成的 gRPC 客户端改为 HTTP 请求形式，以便在浏览器中直接使用。

---

## 🛠 安装与工具链准备

### 安装 TypeScript gRPC 生成工具

使用 [@protobuf-ts/plugin](https://www.npmjs.com/package/@protobuf-ts/plugin)，而**不是**其他 `protoc-gen-ts` 插件：

```bash
npm install -g @protobuf-ts/plugin
````

确认是否安装成功：

```bash
which protoc-gen-ts
```

---

## 🧱 基础开发环境示例

```bash
# 示例1
npm list -g --depth=0
├── @protobuf-ts/plugin@2.9.4
├── ts-node@10.9.2
├── typescript@5.4.5

# 示例2
npm list -g --depth=0
├── @protobuf-ts/plugin@2.9.4
├── vite@5.4.8
```

---

## 📦 在 Kratos 项目中生成 TS 客户端代码

### 示例 Makefile 规则（推荐）

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

> 更多例子见：[demo1 Makefile](https://github.com/orzkratos/vue3kratos-demos/blob/main/demo1kratos/Makefile)

---

## ⚙️ 替换 gRPC 为 HTTP 请求

生成的客户端默认基于 gRPC。如果你希望通过 HTTP 进行调用，可使用下列工具自动转换：

### 安装 CLI 工具

```bash
go install github.com/orzkratos/vue3kratos/cmd/vue3orzkratos@latest
```

### 使用 CLI 替换客户端文件中的请求方式

```bash
vue3orzkratos gen-grpc-via-http-in-path \
  --grpc-ts-path=/absolute/path/to/your.client.ts
```

该命令会将目标文件中 gRPC 的调用方式替换为基于 Axios 的 HTTP 请求。

---

## 📦 安装 Axios HTTP 客户端模块

你还需要在前端项目中安装辅助包：

```bash
npm install @yyle88/grpt
```

> 模块地址：[@yyle88/grpt on npm](https://www.npmjs.com/package/@yyle88/grpt)

---

## 🔁 示例项目

* [demo1kratos](https://github.com/orzkratos/vue3kratos-demos/tree/main/demo1kratos) – Kratos 项目生成 TypeScript 客户端代码
* [demo2kratos](https://github.com/orzkratos/vue3kratos-demos/tree/main/demo2kratos) – HTTP 转换示例

---

## ✅ 特性小结

* 将 Kratos proto 文件生成为 TypeScript gRPC 客户端
* 支持自动替换为 HTTP 请求形式（Axios 封装）
* 支持类型提示与自动补全，前端集成体验极佳
* 可集成至 Makefile 或 CI/CD 中

---

## 💡 适合人群

* 使用 Kratos 作为后端的开发者
* 想通过 Vue 3 编写前端，直接调用后端服务
* 希望实现“全类型安全”的调用体验

---

## 旧版本文档

[旧版说明](internal/docs/README_OLD_DOC.zh.md)

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
