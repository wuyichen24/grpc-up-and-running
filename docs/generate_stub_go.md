## Generate Server Stub
- Install protoc (protocol buffer compiler)
- Install the gRPC library
  ```bash
  go get -u google.golang.org/grpc
  ```
- Install the protoc plug-in for Go
  ```bash
  go get -u github.com/golang/protobuf/protoc-gen-go
  ```
- Make sure `$GOPATH/bin` in the $PATH
  ```bash
  PATH=$PATH:$GOPATH/bin
  ```
- Run protoc
  ```
  protoc -I <path_of_directory_storing_proto_file> <path_of_proto_file> --go_out=plugins=grpc:<path_of_where_you_want_to_generate_stub_file>
  ```
