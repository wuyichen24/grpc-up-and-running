package main

import (
	"context"
	"log"
	"time"

	pb "grpc-up-and-running/examples/security/one-way-tls/client/ecommerce"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
)

var (
	address = "localhost:50051"
	hostname = "localhost"
	crtFile = "server.crt"        // server public certificate.
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call remote methods
	c.AddProduct(ctx, nil)
	c.GetProduct(ctx, nil)
}