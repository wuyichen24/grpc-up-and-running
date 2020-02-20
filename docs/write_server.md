## Write Server Code
- [**Build Listener**]()
- [**Build gRPC Server**]()
- [**Register Service(s)**]()
- [**Start gRPC Server**]()

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

### Register Service(s)
- Register single service
  ```go
  type server struct {}
  
  pb.RegisterAbcServer(s, &server{})    // Abc is the service name
  ```
- Register multiple services
  ```go
  type abcServer struct {}
  type xyzServer struct {}
  
  abc_pb.RegisterAbcServer(s, &abcServer{})   // Abc is the service name
  xyz_pb.RegisterXyzServer(s, &xyzServer{})   // Xyz is the service name
  ```
  
  ### Start gRPC Server
  ```go
  if err := s.Serve(lis); err != nil {
      log.Fatalf("failed to serve: %v", err)
  }
  ```
