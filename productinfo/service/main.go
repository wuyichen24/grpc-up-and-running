package main

import (
	"log"
	"net"

	//pb "grpc-up-and-running/productinfo/service/ecommerce"
	pb "github.com/wuyichen24/grpc-up-and-running/productinfo/service/ecommerce"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
