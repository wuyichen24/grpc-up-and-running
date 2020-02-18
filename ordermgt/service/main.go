package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	ordermgt_pb "ordergmt/service/ecommerce"
	hello_pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(orderUnaryServerInterceptor),     // Register unary interceptor.
		grpc.StreamInterceptor(orderServerStreamInterceptor))   // Register stream interceptor.

	// Register 2 services: OrderManagement and Hello
	// Example of Multiplexing - Run multiple services on one gRPC server
	ordermgt_pb.RegisterOrderManagementServer(s, &orderMgtServer{})
	hello_pb.RegisterGreeterServer(s, &helloServer{})

	log.Printf("Starting gRPC listener on port " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["102"] = ordermgt_pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"},     Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = ordermgt_pb.Order{Id: "103", Items: []string{"Apple Watch S4"},                      Destination: "San Jose, CA",      Price: 400.00}
	orderMap["104"] = ordermgt_pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = ordermgt_pb.Order{Id: "105", Items: []string{"Amazon Echo"},                         Destination: "San Jose, CA",      Price: 30.00}
	orderMap["106"] = ordermgt_pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"},      Destination: "Mountain View, CA", Price: 300.00}
}
