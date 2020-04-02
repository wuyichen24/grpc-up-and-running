package main

import (
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-up-and-running/examples/security/oauth2/client/ecommerce"
)

var (
	address  = "localhost:50051"
	hostname = "localhost"
	crtFile  = "server.crt"        // server public certificate.
)

func main() {
	auth := oauth.NewOauthAccess(fetchToken())

	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
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


func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}