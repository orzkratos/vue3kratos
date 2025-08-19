# vue3kratos

> Vue 3 frontends and Kratos backends integration toolkit  
> Generate type-safe TypeScript clients from Go Kratos proto files  
> Seamless TypeScript + Go RPC integration, use `@protobuf-ts/plugin`  
> Full type safety from backend to frontend  
> Call backend APIs like local functions

## CHINESE README

[中文说明](README.zh.md)

## Core Architecture

`vue3kratos` bridges Go backends and Vue 3 frontends with TypeScript client.

### Development Toolchain
```
+-------------+    +----------+    +---------------+    +--------------+    +---------------+
| .proto files| -> | protoc   | -> | gRPC TS Client| -> | vue3kratos   | -> | HTTP TS Client|
|             |    | + plugin |    |               |    | CLI Convert  |    |               |
+-------------+    +----------+    +---------------+    +--------------+    +---------------+
```

### Runtime Architecture
```
+------------------+                                    +------------------+
|   Vue 3 Frontend |                                    |  Kratos Backend  |
|                  |                                    |                  |
| ┌──────────────┐ |                                    | ┌──────────────┐ |
| │ gRPC客户端代码│ |  Code writes gRPC-style calls      │ │   HTTP Service│ |
| │ client.call() │ |                                    | │  :28000       │ |
| └──────┬───────┘ |                                    | └──────────────┘ |
|        │         |                                    |                  |
| ┌──────▼───────┐ |   Runtime sends HTTP requests      |                  |
| │HTTP Convert   │ |  =========================>        |                  |
| │(@yyle88/grpt) │ |     POST /api/method               |                  |
| └──────────────┘ |     Content-Type: application/json |                  |
|                  |                                    | ┌──────────────┐ |
|                  |                                    | │   gRPC Service│ |
|                  |                    ❌ Not use ---> | │   :28001      │ |
|                  |                                    | │  (Not used)   │ |
|                  |                                    | └──────────────┘ |
+------------------+                                    +------------------+
```

## 🌟 Highlights

*   **Auto Code Generation**: Generate clean TypeScript clients from proto files
*   **No Manual API Writing**: Say goodbye to handwriting API clients
*   **Single Command Convert**: Transform gRPC clients to HTTP with one command
*   **Web Compatible**: Direct use in web without gRPC complexity
*   **Full Type Safety**: Complete type checking from backend to frontend
*   **IDE Autocompletion**: Rich development experience with smart suggestions
*   **Makefile Integration**: Easy integration into existing build process
*   **CI/CD Pipeline**: Smooth workflow automation support
*   **Axios HTTP Clients**: Modern HTTP client implementation
*   **Native Function Feel**: Call APIs like local functions

## 🚀 Quick Start

Experience the magic of `vue3kratos` in minutes.

### 1. Run the Backend Service

1st, run the `demo1go` backend service.

```bash
# Navigate to the backend service DIR
cd internal/demos/demo1x/demo1go

# Start the service
make run
# or
go run ./cmd/demo1go
```

The service will start on `http://127.0.0.1:28000` (HTTP) and `http://127.0.0.1:28001` (gRPC).

### 2. Run the Frontend Demo

Next, in another terminal, run the `vue3npm` frontend demo.

```bash
# Navigate to the frontend demo DIR
cd internal/demos/demo2x/vue3npm

# Install dependencies
npm install

# Run wise demo
npm run demo:wise
```

You can now see real-time logs in the console from the frontend application calling the backend API, showcasing the seamless co-work between them.

## The `vue3kratos` Flow Explained

Using `vue3kratos` in your project involves just a few steps.

### Step 1: Install the Chain

Ensure you have `@protobuf-ts/plugin` installed, and the `vue3kratos` Go CLI app.

```bash
# Install the protobuf-ts plugin (use this one, NOT other protoc-gen-ts plugins)
npm install -g @protobuf-ts/plugin

# Verify the installation
which protoc-gen-ts

# Install the vue3kratos CLI
go install github.com/orzkratos/vue3kratos/cmd/vue3orzkratos@latest
```

#### Example Development Environment

```bash
# Example global dependencies
npm list -g --depth=0
├── @protobuf-ts/plugin@2.9.4
├── typescript@5.4.5
├── vite@5.4.8
```

### Step 2: Generate the gRPC Client

In the Kratos project, use `protoc` and `protoc-gen-ts` to generate the TypeScript client from the `.proto` files. It is recommended to manage this in a `Makefile`.

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

Also add this variable to your Makefile:

```makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
```

### Step 3: Convert to HTTP Client

Use the `vue3orzkratos` CLI to convert the gRPC client file generated in the previous step into an HTTP client.

```bash
vue3orzkratos gen-grpc-via-http-in-path  --grpc-ts-path=/path/to/the/generated.client.ts
```

This command modifies the target file by replacing all gRPC calls with Axios HTTP requests.

### Step 4: Use in Vue

In the Vue project, install the `@yyle88/grpt` helper module. Then you can call APIs as if they are native functions.

```bash
npm install @yyle88/grpt
```

> npm module URL: [@yyle88/grpt](https://www.npmjs.com/package/@yyle88/grpt)

```typescript
// Real demo from internal/demos/demo2x/vue3npm/src/demo-ping.ts
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { RpcpingClient } from "./rpc/rpcping/rpcping.client";
import { StringValue } from "./rpc/google/protobuf/wrappers";

// Create transport instance
const demoTransport = new GrpcWebFetchTransport({
    baseUrl: "http://127.0.0.1:28000",
    meta: {
        Authorization: 'TOKEN-888',
    },
});

const rpcpingClient = new RpcpingClient(demoTransport);

// Call API example
async function demoPing() {
    const request = StringValue.create({
        value: "Hello from Vue3 Kratos!"
    });

    try {
        const response = await rpcpingClient.ping(request, {});
        console.log('Ping success:', response.data.value);
        return response.data.value;
    } catch (err) {
        console.error('Ping FAILED:', err);
        throw err;
    }
}
```

---

## 🔁 Demo Projects

* **[internal/demos/demo1x/demo1go](internal/demos/demo1x/demo1go)** – Generate TypeScript gRPC client from Kratos project
* **[internal/demos/demo2x/vue3npm](internal/demos/demo2x/vue3npm)** – HTTP conversion and Vue 3 integration example

---

## ✅ Feature Summary

* Generate type-safe TypeScript gRPC clients from Kratos proto files
* Support automatic conversion to HTTP requests (Axios-based)
* Full type safety with IDE autocomplete and error checking
* Easy integration into Makefiles or CI/CD pipelines
* Browser-compatible HTTP clients for direct frontend use

---

## 💡 Who Should Use This

* Developers using Kratos as their backend
* Frontend developers using Vue 3 to directly call backend services
* Anyone who wants a fully type-safe client-server integration experience

---

## OLD-Documentation

[README OLD DOC](internal/doc/README_OLD_DOC.en.md)

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!
