# gRPC Gateway

## Background
- Add a reverse proxy server in front of gRPC server to expose RESTful JSON API for each remote method in the gRPC service and accept HTTP requests from REST clients.
- Provide the ability to invoke gRPC service in both gRPC and RESTful ways.

## Installation
- Make sure the Protocol Buffer Compiler has been installed properly.
- Download some dependent packages
  ```bash
  go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
  go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
  go get -u github.com/golang/protobuf/protoc-gen-go
  ```
  
## Modify The Service Definition File (*.proto)
- Example
  ```proto
  syntax = "proto3";

  import "google/protobuf/wrappers.proto";
  import "google/api/annotations.proto";

  package ecommerce;

  service ProductInfo {
      rpc addProduct(Product) returns (google.protobuf.StringValue) {
          option (google.api.http) = {
              post: "/v1/product"
              body: "*"
          };
      }
      rpc getProduct(google.protobuf.StringValue) returns (Product) {
          option (google.api.http) = {
              get:"/v1/product/{value}"
          };
      }
  }

  message Product {
      string id = 1;
      string name = 2;
      string description = 3;
      float price = 4;
  }
  ```
- Rules
   - Each mapping needs to specify a URL path template and an HTTP method.
   - The path template can contain one or more fields in the gRPC request message. But those fields should be nonrepeated fields with primitive types.
   - Any fields in the request message that are not in the path template automatically become HTTP query parameters if there is no HTTP request body.
   - Fields that are mapped to URL query parameters should be either a primitive type or a repeated primitive type or a nonrepeated message type.
   - For a repeated field type in query parameters, the parameter can be repeated in the URL.
     `...?param=A&param=B.`
   - For a message type in query parameters, each field of the message is mapped to a separate parameter.
     `...?foo.a=A&foo.b=B&foo.c=C`

## Generate Service Stub (*.pb.go)
- Change directory to the directory of `main.go` and `go.mod`.
- Run the command
  ```bash
  protoc -I <path_of_directory_storing_proto_file> <path_of_proto_file> \
  -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:<path_of_directory_where_you_want_to_generate_stub_file>
  ```
