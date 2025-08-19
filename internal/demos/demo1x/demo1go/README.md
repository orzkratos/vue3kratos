# Demo1Go - Kratos Backend Service

A simple kratos project (no config. no errors. no bin. no wire. no database).

A simple kratos project provides a minimalist Kratos backend service, designed to work as the server-side counterpart for the `vue3kratos` frontend demonstrations.

## CHINESE README

[中文说明](README.zh.md)

## 🌟 Features

*   **Simple Kratos Setup**: A lightweight project with no complex dependencies like configs, databases.
*   **Double Protocol Support**: Exposes both gRPC (port `28001`) and HTTP (port `28000`) endpoints at once.
*   **Comprehensive Demo APIs**: Implements three distinct services to showcase different API patterns:
    *   **Greeter**: A basic "Hello World" service.
    *   **Ping**: A simple Ping/Pong service.
    *   **RpcDemo**: A full CRUD (Create, Read, Update, Delete) service.
*   **Frontend Support**: Built to work with the `vue3npm` frontend demo project.

## 🚀 Getting Started

### Prerequisites

Ensure you have a working Go environment (Go 1.18+ is recommended).

### Running the Service

1.  Navigate to the project DIR:
    ```bash
    cd /Users/admin/go-projects/orzkratos/vue3kratos/internal/demos/demo1x/demo1go
    ```
2.  Run the application:
    ```bash
    go run ./cmd/demo1go
    ```
3.  You should see the following output, confirming the service is running:
    ```
    INFO msg=[HTTP] server listening on: [::]:28000
    INFO msg=[gRPC] server listening on: [::]:28001
    ```

## 🔬 API Endpoints

You can test the running service using `curl` or any API client.

### Greeter Service

*   **Endpoint**: `GET /api/greeter/meet/{name}`
*   **Description**: A basic greeting service.

```bash
curl "http://127.0.0.1:28000/api/greeter/meet/gogogo"
```

### Ping Service

*   **Endpoint**: `GET /api/service/ping`
*   **Description**: A simple ping/pong test.

```bash
curl "http://127.0.0.1:28000/api/service/ping?value=gogogo"
```

### RpcDemo CRUD Service

This service demonstrates a full set of CRUD operations.

#### 1. Create (POST)
*   **Endpoint**: `/api/demo/create-rpc-demo`
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

#### 2. Delete (DELETE)
*   **Endpoint**: `/api/demo/create-rpc-demo/{code}`
```bash
curl -X DELETE "http://127.0.0.1:28000/api/demo/create-rpc-demo/C002"
```

#### 3. Update (PUT)
*   **Endpoint**: `/api/demo/update-rpc-demo`
```bash
curl -X PUT "http://127.0.0.1:28000/api/demo/update-rpc-demo" \
  -H "Content-Type: application/json" \
  -d '{"code":"C001","name":"N001-V2"}'
```

#### 4. Select (GET)
*   **Endpoint**: `/api/demo/select-rpc-demo`
```bash
curl "http://127.0.0.1:28000/api/demo/select-rpc-demo?type=T001"
```

## Frontend Integration

This backend service works with the `vue3npm` demo:

1. Start this backend service first
2. Navigate to the `vue3npm` DIR  
3. Run demo scripts (`npm run demo:wise`, `npm run demo:epic`)
4. Watch frontend and backend interact
5. Enjoy real-time communication

## Thanks

Thank you! Happy Coding! 🎉 🎉 🎉
