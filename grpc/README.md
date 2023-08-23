## Install protoc
Using protoc to compile proto file and generate code. Open blow website to select myself os version to download and install it
1. https://grpc.io/docs/protoc-installation/
2. https://github.com/protocolbuffers/protobuf/releases/tag/v24.1

## Install Golang Library
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$ go get google.golang.org/protobuf
$ go get google.golang.org/grpc
```

## Write Proto file
blow example from [grpc-go](https://github.com/grpc/grpc-go/blob/master/examples/helloworld)
```protobuf
syntax = "proto3";

option go_package = "api/v1";
option java_multiple_files = true;
option java_package = "io.grpc.examples.helloworld";
option java_outer_classname = "HelloWorldProto";

package activity;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

## Generate Protobuf and GRPC code
```shell
$ cd grpc
$ protoc protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/api/v1/*.proto
```

## Run Code
```shell
# run server
go run server/v1/main.go

# run client
go run client/v1/main.go
```

## Reference
1. [protobuf](https://protobuf.dev/)
2. [grpc-quickstart](https://grpc.io/docs/languages/go/quickstart/)

