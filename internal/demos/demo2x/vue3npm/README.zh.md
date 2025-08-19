# Vue3 Kratos 演示项目

演示 [vue3kratos](https://github.com/orzkratos/vue3kratos) 工具链如何将 Vue 3 前端与 Kratos 后端进行无缝集成。

## 英文文档

[ENGLISH README](README.md)

## 🌟 项目特色

- **gRPC via HTTP**: 将 gRPC 调用透明转换为 HTTP 请求
- **类型安全**: 完整的 TypeScript 类型支持和编译时检查
- **自动生成**: 从 proto 文件自动生成客户端代码
- **浏览器兼容**: 解决 gRPC 在浏览器中的兼容性问题

## 📋 可用演示

### 核心演示脚本

| 脚本命令 | 演示内容 | HTTP方法 | 接口描述 |
|---------|---------|---------|----------|
| `npm run demo:wise` | **全接口演示** | ALL | 🌟 **推荐：所有服务的综合演示** |
| `npm run demo:epic` | 史诗级服务 | GET | 基础问候服务，最简单的演示 |
| `npm run demo:ping` | RpcPing 服务 | GET | Ping/Pong 测试服务 |
| `npm run demo:crud` | RpcDemo CRUD | POST/PUT/GET/DELETE | 完整的增删改查操作演示 |

### 别名脚本

| 别名 | 等效命令 | 说明 |
|------|---------|------|
| `npm run demo:wise` | `npm run demo:wise` | 综合演示 |
| `npm run demo:epic` | `npm run demo:epic` | 史诗级服务演示 |
| `npm run demo:ping` | `npm run demo:ping` | Ping 服务演示 |
| `npm run demo:crud` | `npm run demo:crud` | CRUD 操作演示 |

## 🚀 快速开始

### 1. 前置条件

确保已安装依赖：
```bash
npm install
```

### 2. 启动后端服务

```bash
# 进入后端服务目录
cd ../demo1go

# 启动服务
make run
# 或者
go run cmd/demo1go/main.go
```

服务将在以下端口启动：
- **HTTP服务**: `http://127.0.0.1:28000`
- **gRPC服务**: `http://127.0.0.1:28001`

### 3. 运行演示

```bash
# 🌟 推荐：运行完整演示
npm run demo:wise

# 或运行单个演示
npm run demo:epic        # 史诗级服务
npm run demo:ping        # RpcPing 服务
npm run demo:crud        # RpcDemo CRUD
```

## 📋 演示详情

### Demo1: 智慧演示 (`demo-wise.ts`)
- 按顺序调用所有服务接口
- 提供详细的执行状态反馈
- 生成完整的成功率统计报告
- 最佳的整体功能验证方案

### Demo2: 史诗级服务 (`demo-epic.ts`)
- **接口**: `SayHello`
- **方法**: GET `/api/greeter/meet/{name}`
- **功能**: 基础的问候服务
- **特点**: 最简单的 gRPC-via-HTTP 示例

### Demo3: RpcPing 服务 (`demo-ping.ts`)
- **接口**: `Ping`
- **方法**: GET `/api/service/ping`
- **功能**: Ping/Pong 测试
- **特点**: 使用 `google.protobuf.StringValue` 类型

### Demo4: RpcDemo CRUD (`demo-crud.ts`)
- 按顺序调用所有服务接口
- 提供详细的执行状态反馈
- 生成完整的成功率统计报告
- 最佳的整体功能验证方案

## 🛠 技术架构

### 核心技术栈
- **前端框架**: Vue 3 + TypeScript
- **RPC通信**: gRPC-Web + HTTP 适配
- **协议定义**: Protocol Buffers
- **包管理**: npm + ES Modules
- **类型系统**: 完整的 TypeScript 类型安全

### 关键组件
- **[@protobuf-ts/plugin](https://npmjs.com/package/@protobuf-ts/plugin)**: TypeScript gRPC 代码生成
- **[@yyle88/grpt](https://npmjs.com/package/@yyle88/grpt)**: gRPC-to-HTTP 适配器
- **vue3kratos CLI**: 自动化代码转换工具

## ⚙️ 配置说明

### 服务端配置
修改后端服务地址（如需要）：
```typescript
const demoTransport = new GrpcWebFetchTransport({
    baseUrl: "http://127.0.0.1:28000",  // 后端服务地址
    meta: {
        Authorization: 'TOKEN-888',      // 认证令牌
    },
});
```

### 项目配置要求
- **package.json**: 必须设置 `"type": "module"`
- **tsconfig.json**: 建议设置 `"target": "ES2023"` 以支持 BigInt

## 🔧 开发指南

### 添加新的演示
1. 在 `src/` 目录创建新的 `.ts` 文件
2. 导入相应的客户端和消息类型
3. 编写演示逻辑和错误处理
4. 在 `package.json` 的 scripts 中添加对应命令

### 自定义客户端使用
```typescript
import { RpcpingClient } from "./rpc/rpcping/rpcping.client";
import { StringValue } from "./rpc/google/protobuf/wrappers";

const client = new RpcpingClient(demoTransport);
const request = StringValue.create({ value: "Hello from Vue3 Kratos!" });

try {
    const response = await client.ping(request, {});
    console.log('成功:', response.data.value);
} catch (err) {
    console.error('失败:', err);
}
```

## 📝 重要说明

### 依赖关系
- **后端服务**: 必须先启动 demo1go 后端服务
- **网络连接**: 确保前后端网络畅通
- **端口占用**: 默认使用 28000/28001 端口
- **Proto一致性**: 确保前后端 proto 定义完全一致

### 最佳实践
- 推荐使用 `npm run demo:wise` 进行完整功能验证
- 开发时先运行单个演示进行调试
- 生产环境需要相应调整服务地址和认证配置
- 定期检查依赖包版本更新

## 🎯 学习价值

通过本演示项目，您将学习到：

1. **vue3kratos 工具链的完整使用流程**
2. **gRPC 与 HTTP 的透明转换原理和实践**
3. **TypeScript 与 Protocol Buffers 的深度集成**
4. **现代前端与微服务后端的最佳架构实践**
5. **类型安全的 API 调用模式**

## 🔧 传统支持

该项目也支持用于教育目的的传统方法：

### 使用 copyrpc 主函数
通过外部 `copyrpc` 主函数执行以获得所需代码。

### 使用 grpcrewrite 命令
使用 JavaScript 命令转换生成的客户端代码：

```bash
# 本地路径示例
npm run grpcrewrite -- src/rpc/v1/greeter/greeter.client.ts

# 绝对路径（推荐实际使用）
npm run grpcrewrite -- /absolute/path/to/your/client.ts
```

生成的代码与 `copyrpc` 输出完全等效。

### 配置说明
- package.json 需要设置 `"type": "module"` 以支持 JavaScript 模块
- tsconfig.json 需要设置 `"target": "ES2023"` 以支持 BigInt

## 🔗 相关链接

- [vue3kratos 主项目](https://github.com/orzkratos/vue3kratos)
- [Kratos 框架](https://go-kratos.dev/)
- [@protobuf-ts/plugin](https://npmjs.com/package/@protobuf-ts/plugin)
- [@yyle88/grpt](https://npmjs.com/package/@yyle88/grpt)

---

💡 **提示**: 建议首先运行 `npm run demo:wise` 来体验完整的功能演示！