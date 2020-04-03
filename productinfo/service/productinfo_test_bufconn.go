// Test for AddProduct by bufconn
// bufconn can avoid the server to open up a port the client connects to.
package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	pb "grpc-gateway/server/ecommerce"
	"testing"
	"time"
)

const (
	bufSize = 1024 * 1024
)

var listener *bufconn.Listener

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

// Initialization of BufConn.
// Package bufconn provides a net.Conn implemented by a buffer and related dialing and listening functionality.
func initGRPCServerBuffConn() {
	listener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	// Register reflection server on gRPC server.
	reflection.Register(s)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}

// Test AddProduct using Buffconn
func TestServer_AddProductBufConn(t *testing.T) {
	ctx := context.Background()
	initGRPCServerBuffConn()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	// Contact the server and print out its response.
	name := "Sumsung S10"
	description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
	price := float32(700.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf(r.Value)
}
