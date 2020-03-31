## Write Server Code

- [**Implement Main Method**](#implement-main-method)
   - [**Build Listener**](#build-listener)
   - [**Build gRPC Server**](#build-grpc-server)
   - [**Available Server Option**](#available-server-option)
   - [**Register Service(s)**](#register-services)
   - [**Start gRPC Server**](#start-grpc-server)
- [**Implement Remote Methods**](#implement-remote-methods)
   - [Basic Pattern](#basic-pattern)
   - [Process Inputs](#process-inputs)
      - Process unary input
      - Process stream intput
      - Process metadata
   - [Return Outputs](#return-outputs)
      - Return unary output
      - Return stream output
      - Return metadata
      - Return error

## Implement Main Method
### Build Listener
```go
listener, err := net.Listen("tcp", ":50051")
```

### Build gRPC Server
- Basic version
  ```go
  s := grpc.NewServer()
  ```
- With server option(s)
  ```go
  opts := []grpc.ServerOption{
      opt1, opt2, opt3
  }
  s := grpc.NewServer(opts...)
  ```
  
### Available Server Option
- Unary interceptor
  ```go
  grpc.UnaryInterceptor(unaryInterceptorFunc)
  ```
- Stream interceptor
  ```go
  grpc.StreamInterceptor(streamInterceptorFunc)
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

## Implement Remote Methods
### Basic Pattern
```go
type abcServer struct {}

func (s *abcServer) RemoteFunc1(ctx context.Context, input *InputType) (*OutputType, error) {}

func (s *abcServer) RemoteFunc2(ctx context.Context, input *InputType) (*OutputType, error) {}

func (s *abcServer) RemoteFunc3(ctx context.Context, input *InputType) (*OutputType, error) {}
```

### Process Inputs
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

### Return Outputs
- Return unary output
  ```go
  func (s *abcServer) RemoteFunc(ctx context.Context, input *InputType) (*OutputType, error) {
      return &output, status.New(codes.OK, "").Err()
  }
  ```
- Return stream output
  ```go
  func (s *abcServer) RemoteFunc(input *InputType, stream pb.Abc_RemoteFuncServer) error {
      for _ , output := range output {
          err := stream.Send(&output)
      }
      return nil
  }
  ```
- Return metadata
   - In unary input function
     ```go
     func (s *abcServer) RemoteFunc(ctx context.Context, input *InputType) (*OutputType, error) {
         header := metadata.Pairs("header-key", "val")
         grpc.SendHeader(ctx, header)
         trailer := metadata.Pairs("trailer-key", "val")
         grpc.SetTrailer(ctx, trailer)
     }
     ```
   - In stream input function
     ```go
     func (s *abcServer) RemoteFunc(stream pb.Abc_RemoteFuncServer) error {
         header := metadata.Pairs("header-key", "val")
         stream.SendHeader(header)
         trailer := metadata.Pairs("trailer-key", "val")    
         stream.SetTrailer(trailer)
     }
     ```
- Return error
   - Only error status
     ```go
     errorStatus := status.New(codes.ErrorCodeOption, "The error description.")  // ErrorCodeOption needs to be replaced by real option.
     return errorStatus.Err()
     ```
   - error status with details
     ```go
     import epb "google.golang.org/genproto/googleapis/rpc/errdetails"
     
     errorStatus := status.New(codes.ErrorCodeOption, "The error description.")  // ErrorCodeOption needs to be replaced by real option.
     ds, err := errorStatus.WithDetails(
			   &epb.DetailOption1{},                                                   // DetailOption1 needs to be replaced by real option.
         &epb.DetailOption2{},                                                   // DetailOption2 needs to be replaced by real option.
         &epb.DetailOption3{}                                                    // DetailOption3 needs to be replaced by real option.
		 )
     return ds.Err()
     ```

