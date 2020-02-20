## Write Server Code
- [**Build Listener**]()
- [**Build gRPC Server**]()

### Build Listener
```go
listener, err := net.Listen("tcp", ":50051")
```

### Build gRPC Server
- Basic version
  ```go
  s := grpc.NewServer()
  ```
- With unary interceptor
  ```go
  s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptorFunc))
  ```
- With stream interceptor
  ```go
  s := grpc.NewServer(grpc.StreamInterceptor(streamInterceptorFunc))
  ```
- With both unary interceptor and stream interceptor
  ```go
  s := grpc.NewServer(
    grpc.UnaryInterceptor(unaryInterceptorFunc),     // Register unary interceptor.
    grpc.StreamInterceptor(streamInterceptorFunc))   // Register stream interceptor.
  ```
