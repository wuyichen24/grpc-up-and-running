## Write Server Code
- [**Build Listener**]()
- [**Build gRPC Server**]()
- [**Register Service(s)**]()
- [**Start gRPC Server**]()
- [**Implement Remote Methods**]()

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
  type abcServer struct {}
  
  pb.RegisterAbcServer(s, &abcServer{})       // Abc is the service name
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

### Implement Remote Methods
#### Basic Pattern
```go
type abcServer struct {}

func (s *abcServer) RemoteFunc1(ctx context.Context, input *InputType) (*OutputType, error) {}

func (s *abcServer) RemoteFunc2(ctx context.Context, input *InputType) (*OutputType, error) {}

func (s *abcServer) RemoteFunc3(ctx context.Context, input *InputType) (*OutputType, error) {}
```

#### Process Inputs
- Process unary input
  ```go
  type abcServer struct {}
  
  func (s *abcServer) RemoteFunc(ctx context.Context, input *InputType) (*OutputType, error) {
      // Use input directly.
  }
  ```
- Process stream intput
  ```go
  type abcServer struct {}
  
  func (s *abcServer) RemoteFunc(stream pb.Abc_RemoteFuncServer) error {
      for {
          input, err := stream.Recv()
          if err == io.EOF {
              return stream.SendAndClose(output)       // At the end of input stream, return the output.
          }
          
          if err != nil {
              return err
          }
          
          // Process single input.
      }
  }
  ```
- Process metadata
   - In unary input function
     ```go
     func (s *abcServer) RemoteFunc(ctx context.Context, input *InputType) (*OutputType, error) {
         md, ok := metadata.FromIncomingContext(ctx)
     }
     ```
   - In stream input function
     ```go
     func (s *abcServer) RemoteFunc(stream pb.Abc_RemoteFuncServer) error {
         md, ok := metadata.FromIncomingContext(stream.Context())
     }
     ```

#### Return Outputs
- Return unary output
- Return stream output
- Return error
- Return metadata
