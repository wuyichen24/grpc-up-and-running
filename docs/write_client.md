## Write Client Code
- [**Build Connection**]()
- [**Build Client**]()
- [**Build Context**]()
- [**Create Metadata (Optional)**]()
- [**Add Metadata to Context (Optional)**]()

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
- With metadata
  ```go
  import "google.golang.org/grpc/metadata"
  ctx := metadata.NewOutgoingContext(context.Background(), md)
  ```

### Create Metadata (Optional)
- Option 1
  ```go
  md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})
  ```
- Option 2
  ```go
  md := metadata.Pairs(
    "key1", "val1",
    "key2", "val2"
  )
  ```

### Add Metadata to Context (Optional)
- Create a new context by metadata
  ```go
  ctx := metadata.NewOutgoingContext(context.Background(), md)
  ```
- Append metadata to an existing context
  ```go
  ctx2 := metadata.AppendToOutgoingContext(ctx1, "key1", "val1", "key2", "val2")
  ```
  

