GOHOSTOS:=$(shell go env GOHOSTOS)

ifeq ($(GOHOSTOS), windows)
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
	THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell $(Git_Bash) -c "find proto3ps/google/api -name *.proto")
else
	API_PROTO_FILES=$(shell find api -name *.proto)
	THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find proto3ps/google/api -name *.proto)
endif

.PHONY: api
# generate internal proto
api:
	protoc --proto_path=./api \
		   --proto_path=./proto3ps \
		   --go_out=paths=source_relative:./api \
		   --go-http_out=paths=source_relative:./api \
		   --go-grpc_out=paths=source_relative:./api \
		   --openapi_out=fq_schema_naming=true,default_response=false,title=DEMO-TITLE:. \
		   $(API_PROTO_FILES)

.PHONY: orz
# generate internal proto
orz:
	make api
	# go install github.com/orzkratos/orzkratos/cmd/orzkratos-srv-proto@latest
	orzkratos-srv-proto -auto
	go mod tidy -e
	echo "success"

# 这就是新增的内容
.PHONY: web-api-grpc-ts
# npm install -g @protobuf-ts/plugin
web-api-grpc-ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	PROTOC_GEN_TS=$$(which protoc-gen-ts) && \
	protoc \
	--plugin=protoc-gen-ts=$$PROTOC_GEN_TS \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./api \
	--proto_path=./proto3ps \
	$(API_PROTO_FILES)

	PROTOC_GEN_TS=$$(which protoc-gen-ts) && \
	protoc \
	--plugin=protoc-gen-ts=$$PROTOC_GEN_TS \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./proto3ps \
	$(THIRD_PARTY_GOOGLE_API_PROTO_FILES)

# 这就是新增的内容
.PHONY: web-api-grpt-ts
# go install github.com/orzkratos/vue3kratos/cmd/vue3orzkratos@latest
web-api-grpt-ts:
	vue3orzkratos gen-grpc-via-http-in-root --grpc-ts-root=./bin/web_api_grpc_ts.out

# 这就是新增的内容
.PHONY: web-api-cleanup
# cleanup
web-api-cleanup:
	rm -rf ./bin/web_api_grpc_ts.out
