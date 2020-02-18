package main

import (
	"context"
	"log"
	hello_pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

type helloServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *helloServer) SayHello(ctx context.Context, in *hello_pb.HelloRequest) (*hello_pb.HelloReply, error) {
	log.Printf("Greeter Service - SayHello RPC")
	return &hello_pb.HelloReply{Message: "Hello " + in.Name}, nil
}
