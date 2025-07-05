# vue3kratos

> Vue 3 frontends and Kratos backends integration toolkit
> Seamless TypeScript + Go RPC integration, powered by `@protobuf-ts/plugin` and Kratos

---

## CHINESE README

[中文说明](README.zh.md)

---

## ✨ Project Overview

`vue3kratos` is a toolchain that automatically generates TypeScript client code for Vue 3 from [Kratos](https://go-kratos.dev/) backend service definitions (proto files).  
It also supports converting the generated gRPC clients to HTTP-based clients, making them usable directly in browser environments.

---

## 🛠 Installation & Toolchain Setup

### Install TypeScript gRPC Code Generator

Use [@protobuf-ts/plugin](https://www.npmjs.com/package/@protobuf-ts/plugin), **not** any other `protoc-gen-ts` plugins:

```bash
npm install -g @protobuf-ts/plugin
````

Verify the installation:

```bash
which protoc-gen-ts
```

---

## 🧱 Example Development Environments

```bash
# Example 1
npm list -g --depth=0
├── @protobuf-ts/plugin@2.9.4
├── ts-node@10.9.2
├── typescript@5.4.5

# Example 2
npm list -g --depth=0
├── @protobuf-ts/plugin@2.9.4
├── vite@5.4.8
```

---

## 📦 Generate TypeScript Client Code in a Kratos Project

### Recommended Makefile Rule Example

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

Additionally, in your Makefile:

```makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
```

> For more examples, see:
> [demo1 Makefile](https://github.com/orzkratos/vue3kratos-demos/blob/main/demo1kratos/Makefile)

---

## ⚙️ Replace gRPC with HTTP Requests

By default, the generated client is based on gRPC.
If you'd rather use HTTP requests (e.g., in the browser), use the CLI tool below to convert gRPC calls to Axios-based HTTP requests.

### Install CLI Tool

```bash
go install github.com/orzkratos/vue3kratos/cmd/vue3orzkratos@latest
```

### Use CLI to Convert Client File

```bash
vue3orzkratos gen-grpc-via-http-in-path \
  --grpc-ts-path=/absolute/path/to/your.client.ts
```

This command modifies the target file by replacing all gRPC calls with Axios HTTP requests.

---

## 📦 Install Axios-Based HTTP Client Module

Install the helper module in your Vue project:

```bash
npm install @yyle88/grpt
```

> Module URL: [@yyle88/grpt on npm](https://www.npmjs.com/package/@yyle88/grpt)

---

## 🔁 Demo Projects

* [demo1kratos](https://github.com/orzkratos/vue3kratos-demos/tree/main/demo1kratos) – Generate TypeScript gRPC client from Kratos project
* [demo2kratos](https://github.com/orzkratos/vue3kratos-demos/tree/main/demo2kratos) – HTTP conversion example

---

## ✅ Feature Summary

* Generate TypeScript gRPC clients from Kratos proto files
* Support automatic conversion to HTTP requests (Axios-based)
* Full type support with IDE autocomplete
* Easy integration into Makefiles or CI/CD pipelines

---

## 💡 Who Should Use This

* Developers using Kratos as their backend
* Frontend developers using Vue 3 to directly call backend services
* Anyone who wants a fully type-safe client-server integration experience

---

## OLD-Documentation

[README OLD DOC](internal/docs/README_OLD_DOC.en.md)

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
