package main

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-up-and-running/examples/security/one-way-tls/server/ecommerce"
	"log"
	"net"
)

type server struct {}

var (
	port = ":50051"
	crtFile = "server.crt"    // server public certificate.
	keyFile = "server.key"    // server private key.
)

func main() {
	cert, err := tls.LoadX509KeyPair(crtFile,keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterProductInfoServer(s, &server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s server) AddProduct(context.Context, *pb.Product) (*pb.ProductID, error) {
	panic("implement me")
}

func (s server) GetProduct(context.Context, *pb.ProductID) (*pb.Product, error) {
	panic("implement me")
}