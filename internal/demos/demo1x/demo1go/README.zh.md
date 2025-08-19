# Demo1Go - Kratos 后端服务

A simple kratos project (no config. no errors. no bin. no wire. no database).

这是一个极简的 Kratos 后端服务，旨在作为 `vue3kratos` 前端演示的服务器端。

## 英文文档

[ENGLISH README](README.md)

## 🌟 项目特性

*   **简单 Kratos 设置**: 一个轻量级项目，没有复杂的依赖如配置、数据库、Wire。
*   **双协议支持**: 同时暴露 gRPC (端口 `28001`) 和 HTTP (端口 `28000`) 端点。
*   **完备的演示API**: 实现了三个独立的服务以展示不同的API模式：
    *   **Greeter**: 基础的 "Hello World" 服务。
    *   **Ping**: 简单的 Ping/Pong 服务。
    *   **RpcDemo**: 完整的 CRUD (增删改查) 服务。
*   **前端支持**: 为 `vue3npm` 前端演示项目而构建。

## 🚀 快速上手

### 环境要求

请确保你有一个可用的 Go 环境 (推荐 Go 1.18+)。

### 运行服务

1.  进入项目DIR：
    ```bash
    cd /Users/admin/go-projects/orzkratos/vue3kratos/internal/demos/demo1x/demo1go
    ```
2.  运行应用：
    ```bash
    go run ./cmd/demo1go
    ```
3.  你应该能看到以下输出，确认服务已成功运行：
    ```
    INFO msg=[HTTP] server listening on: [::]:28000
    INFO msg=[gRPC] server listening on: [::]:28001
    ```

## 🔬 API 端点

你可以使用 `curl` 或任何 API 客户端来测试运行中的服务。

### Greeter 服务

*   **端点**: `GET /api/greeter/meet/{name}`
*   **描述**: 一个基础的问候服务。

```bash
curl "http://127.0.0.1:28000/api/greeter/meet/gogogo"
```

### Ping 服务

*   **端点**: `GET /api/service/ping`
*   **描述**: 一个简单的 ping/pong 测试。

```bash
curl "http://127.0.0.1:28000/api/service/ping?value=gogogo"
```

### RpcDemo CRUD 服务

这个服务演示了一整套 CRUD 操作。

#### 1. 创建 (POST)
*   **端点**: `/api/demo/create-rpc-demo`
```bash
curl -X POST "http://127.0.0.1:28000/api/demo/create-rpc-demo" \
  -H "Content-Type: application/json" \
  -d '{"code":"C001","name":"N001","type":"T001"}'
```

```bash
curl -X POST "http://127.0.0.1:28000/api/demo/create-rpc-demo" \
  -H "Content-Type: application/json" \
  -d '{"code":"C002","name":"N002","type":"T001"}'
```

#### 2. 删除 (DELETE)
*   **端点**: `/api/demo/create-rpc-demo/{code}`
```bash
curl -X DELETE "http://127.0.0.1:28000/api/demo/create-rpc-demo/C002"
```

#### 3. 更新 (PUT)
*   **端点**: `/api/demo/update-rpc-demo`
```bash
curl -X PUT "http://127.0.0.1:28000/api/demo/update-rpc-demo" \
  -H "Content-Type: application/json" \
  -d '{"code":"C001","name":"N001-V2"}'
```

#### 4. 查询 (GET)
*   **端点**: `/api/demo/select-rpc-demo`
```bash
curl "http://127.0.0.1:28000/api/demo/select-rpc-demo?type=T001"
```

## 前端集成

该后端服务与 `vue3npm` 演示配合使用：

1. 首先启动这个后端服务
2. 进入 `vue3npm` DIR
3. 运行演示脚本 (`npm run demo:wise`, `npm run demo:epic`)
4. 观察前后端交互
5. 享受实时通信

## 谢谢

谢谢您！编程快乐！ 🎉 🎉 🎉
