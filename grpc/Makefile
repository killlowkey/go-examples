.PHONY: generate-v1
generate-v1:
	protoc --go_out=. --go-grpc_out=. protos/api/v1/*.proto

.PHONY: run-server-v1
run-server-v1:
	go run server/v1/main.go

.PHONY: run-client-v1
run-client-v1:
	go run client/v1/main.go

.PHONY: run-server-v2
run-server-v2:
	go run server/v2/main.go

.PHONY: run-client-v2
run-client-v2:
	go run client/v2/main.go

.PHONY: generate-v2
generate-v2:
	protoc --go_out=. --go-grpc_out=. protos/api/v2/*.proto

