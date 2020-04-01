package main

import (
	"context"
	"encoding/base64"
	"log"
	"time"

	pb "grpc-up-and-running/examples/security/basic-auth/client/ecommerce"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
)

var (
	address  = "localhost:50051"
	hostname = "localhost"
	crtFile  = "server.crt"        // server public certificate.
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	auth := basicAuth{
		username: "admin",
		password: "admin",
	}

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(auth),
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

// The struct to hold the collection on fields you want to inject in your RPC calls.
type basicAuth struct {
	username string
	password string
}

// Convert user credentials to request metadata.
func (b basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := b.username + ":" + b.password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "Basic " + enc,
	}, nil
}

// Specify whether channel security is required to pass these credentials.
func (b basicAuth) RequireTransportSecurity() bool {
	return true
}