## Write Client Code

- [Build Connection]()
- [Build Client]()
- [Build Context]()

### Build Connection
- Basic version
  ```go
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  ```
- With unary interceptor
  ```go
  conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryInterceptorFunc))
  ```
- With stream interceptor
  ```go
  conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithStreamInterceptor(streamInterceptorFunc))
  ```
- With both unary interceptor and stream interceptor
  ```go
  conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryInterceptorFunc), grpc.WithStreamInterceptor(streamInterceptorFunc))
  ```

### Build Client
- Basic version
  ```go
  client := pb.NewAbcClient(conn)      // Abc is the service name
  ```

### Build Context
- With timeout
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
  ```
- With deadline
  ```go
  deadline := time.Now().Add(time.Duration(5 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
  ```
