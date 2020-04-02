package main

import (
	"context"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-up-and-running/examples/security/jwt/client/ecommerce"
)

var (
	address      = "localhost:50051"
	hostname     = "localhost"
	crtFile      = "server.crt"        // server public certificate.
	jwtTokenFile = "token.json"        // JWT token file.
)

func main() {
	jwtCreds, err := oauth.NewJWTAccessFromFile(jwtTokenFile)

	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(jwtCreds),
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