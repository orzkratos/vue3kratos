# Vue3 Kratos Demo Project

This project demonstrates the complete [vue3kratos](https://github.com/orzkratos/vue3kratos) toolchain, showcasing seamless integration between Vue 3 frontend and Kratos backend services.

## CHINESE README

[中文说明](README.zh.md)

## 🌟 Project Features

- **gRPC via HTTP**: Transparent conversion of gRPC calls to HTTP requests
- **Type Safety**: Full TypeScript type support with compile-time checking
- **Auto Generation**: Automatic client code generation from proto files
- **Browser Compatible**: Solves gRPC browser compatibility issues

## 📋 Available Demos

### Core Demo Scripts

| Script Command | Demo Content | HTTP Methods | Interface Description |
|---------------|--------------|-------------|---------------------|
| `npm run demo:wise` | **All APIs Demo** | ALL | 🌟 **Recommended: Comprehensive demo of all services** |
| `npm run demo:epic` | Epic Service | GET | Basic greeting service, simplest demo |
| `npm run demo:ping` | RpcPing Service | GET | Ping/Pong test service |
| `npm run demo:crud` | RpcDemo CRUD | POST/PUT/GET/DELETE | Complete CRUD operations demo |

### Alias Scripts

| Alias | Equivalent Command | Description |
|-------|-------------------|-------------|
| `npm run demo:wise` | `npm run demo:wise` | Comprehensive demo |
| `npm run demo:epic` | `npm run demo:epic` | Epic service demo |
| `npm run demo:ping` | `npm run demo:ping` | Ping service demo |
| `npm run demo:crud` | `npm run demo:crud` | CRUD operations demo |

## 🚀 Quick Start

### 1. Prerequisites

Ensure dependencies are installed:
```bash
npm install
```

### 2. Start Backend Services

```bash
# Navigate to backend service directory
cd ../demo1go

# Start services
make run
# or
go run cmd/demo1go/main.go
```

Services will start on these ports:
- **HTTP Service**: `http://127.0.0.1:28000`
- **gRPC Service**: `http://127.0.0.1:28001`

### 3. Run Demos

```bash
# 🌟 Recommended: Run complete demo
npm run demo:wise

# Or run individual demos
npm run demo:epic        # Epic service
npm run demo:ping        # RpcPing service
npm run demo:crud        # RpcDemo CRUD
```

## 📋 Demo Details

### Demo1: Wise Service (`demo-wise.ts`)
- Sequential calls to all service interfaces
- Detailed execution status feedback
- Complete success rate statistics report
- Best overall functionality verification solution

### Demo2: Epic Service (`demo-epic.ts`)
- **Interface**: `SayHello`
- **Method**: GET `/api/greeter/meet/{name}`
- **Function**: Basic greeting service
- **Feature**: Simplest gRPC-via-HTTP example

### Demo3: RpcPing Service (`demo-ping.ts`)
- **Interface**: `Ping`
- **Method**: GET `/api/service/ping`
- **Function**: Ping/Pong testing
- **Feature**: Uses `google.protobuf.StringValue` type

### Demo4: RpcDemo CRUD (`demo-crud.ts`)
- Sequential calls to all service interfaces
- Detailed execution status feedback
- Complete success rate statistics report
- Best overall functionality verification solution

## 🛠 Technical Architecture

### Core Technology Stack
- **Frontend Framework**: Vue 3 + TypeScript
- **RPC Communication**: gRPC-Web + HTTP Adapter
- **Protocol Definition**: Protocol Buffers
- **Package Management**: npm + ES Modules
- **Type System**: Complete TypeScript type safety

### Key Components
- **[@protobuf-ts/plugin](https://npmjs.com/package/@protobuf-ts/plugin)**: TypeScript gRPC code generation
- **[@yyle88/grpt](https://npmjs.com/package/@yyle88/grpt)**: gRPC-to-HTTP adapter
- **vue3kratos CLI**: Automated code transformation tool

## ⚙️ Configuration

### Server Configuration
Modify backend service address (if needed):
```typescript
const demoTransport = new GrpcWebFetchTransport({
    baseUrl: "http://127.0.0.1:28000",  // Backend service address
    meta: {
        Authorization: 'TOKEN-888',      // Authentication token
    },
});
```

### Project Configuration Requirements
- **package.json**: Must set `"type": "module"`
- **tsconfig.json**: Recommend setting `"target": "ES2023"` for BigInt support

## 🔧 Development Guide

### Adding New Demos
1. Create new `.ts` file in `src/` directory
2. Import corresponding client and message types
3. Write demo logic and error handling
4. Add corresponding command in `package.json` scripts

### Custom Client Usage
```typescript
import { RpcpingClient } from "./rpc/rpcping/rpcping.client";
import { StringValue } from "./rpc/google/protobuf/wrappers";

const client = new RpcpingClient(demoTransport);
const request = StringValue.create({ value: "Hello from Vue3 Kratos!" });

try {
    const response = await client.ping(request, {});
    console.log('Success:', response.data.value);
} catch (err) {
    console.error('FAILED:', err);
}
```

## 📝 Important Notes

### Dependencies
- **Backend Service**: Must start demo1go backend service first
- **Network Connection**: Ensure frontend-backend network connectivity
- **Port Usage**: Uses ports 28000/28001 by default
- **Proto Consistency**: Ensure frontend-backend proto definitions are identical

### Best Practices
- Recommend using `npm run demo-all` for complete functionality verification
- Run individual demos first for debugging during development
- Adjust service addresses and authentication configs for production
- Regularly check dependency package version updates

## 🎯 Learning Value

Through this demo project, you will learn:

1. **Complete workflow of vue3kratos toolchain**
2. **Principles and practices of gRPC to HTTP transparent conversion**
3. **Deep integration of TypeScript with Protocol Buffers**
4. **Best architecture practices for modern frontend with microservice backend**
5. **Type-safe API calling patterns**

## 🔧 Legacy Support

This project also supports alternative approaches for educational purposes:

### Using copyrpc main Function
Execute via the outer `copyrpc` main function to get desired code.

### Using grpcrewrite Command
Transform generated client code using JavaScript commands:

```bash
# Local path example
npm run grpcrewrite -- src/rpc/v1/greeter/greeter.client.ts

# Absolute path (recommended for actual use)
npm run grpcrewrite -- /absolute/path/to/your/client.ts
```

The generated code is completely equivalent to `copyrpc` output.

### Configuration Notes
- Requires `"type": "module"` in package.json for JavaScript modules
- Requires `"target": "ES2023"` in tsconfig.json for BigInt support

## 🔗 Related Links

- [vue3kratos Main Project](https://github.com/orzkratos/vue3kratos)
- [Kratos Framework](https://go-kratos.dev/)
- [@protobuf-ts/plugin](https://npmjs.com/package/@protobuf-ts/plugin)
- [@yyle88/grpt](https://npmjs.com/package/@yyle88/grpt)

---

💡 **Tip**: We recommend starting with `npm run demo-all` to experience the complete functionality demonstration!