# Simple Kratos Project

A simple kratos project (no config. no errors. no bin. no wire. no database).

## Run

```bash
go run cmd/demo1go/main.go
```

output log:

```
INFO msg=[HTTP] server listening on: [::]:28000
INFO msg=[gRPC] server listening on: [::]:28001
```

Check apis:

```bash
curl "http://127.0.0.1:28000/api/service/ping?value=gogogo"
```

```bash
curl "http://127.0.0.1:28000/api/greeter/meet/gogogo"
```

Now service provide apis. You can run vue3 demo project to invoke these apis.

## API

### 1. `CreateRpcDemo`（POST + JSON body）

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

---

### 2. `DeleteRpcDemo`（DELETE + 路径参数）

```bash
curl -X DELETE "http://127.0.0.1:28000/api/demo/create-rpc-demo/C002"
```

---

### 3. `UpdateRpcDemo`（PUT + JSON body）

```bash
curl -X PUT "http://127.0.0.1:28000/api/demo/update-rpc-demo" \
  -H "Content-Type: application/json" \
  -d '{"code":"C001","name":"N001-V2"}'
```

---

### 4. `SelectRpcDemo`（GET + query 参数）

```bash
curl "http://127.0.0.1:28000/api/demo/select-rpc-demo?type=T001"
```

## Vue

Now service provide apis. You can run vue3 demo project to invoke these apis.
