## Write Client Code
- [**Build Connection**](#build-connection)
- [**Build Client**](#build-client)
- [**Build Context**](#build-context)
- [**Create Metadata (Optional)**](#create-metadata-optional)
- [**Add Metadata to Context (Optional)**](#add-metadata-to-context-optional)
- [**Call Remote Method**](#call-remote-method)
- [**Handle Response Error**](#handle-response-error)

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
  
### Call Remote Method
- Unary
  ```go
  output, err = client.someRemoteFunc(ctx, &input)
  ```
- Server-side streaming
  ```go
  outputStream, streamErr := client.someRemoteFunc(ctx, input)
  
  for {                                            // Process multiple outputs
      output, err := outputStream.Recv()
      if err == io.EOF {                           // End of stream, break infinite loop
          log.Print("EOF")
          break
      }

      if err == nil {
          // Process output.
      }
  }
  ```
- Client-side streaming
  ```go
  inputStream, err := client.someRemoteFunc(ctx)   // Create input stream
  
  for _ , input := range inputs {                  // Send multiple inputs
      inputStream.Send(&input)
  }
  
  output, err := inputStream.CloseAndRecv()        // Close sending and get output
  ```

### Handle Response Error
```go
import "google.golang.org/grpc/status"

output, err = client.someRemoteFunc(ctx, &input)
if err != nil {
    errorCode    := status.Code(err)
    errorStatus  := status.Convert(err)
    errorDetails := errorStatus.Details()
    
    // Error handling
}
```
