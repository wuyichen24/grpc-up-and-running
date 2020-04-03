package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "examples/grpc-gateway/reverse-proxy/ecommerce"
)

var (
	grpcServerEndpoint = "localhost:50051"
	reverseProxyPort = "8081"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterProductInfoHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)

	if err != nil {
		log.Fatalf("Fail to register gRPC gateway server endpoint: %v", err)
	}

	log.Printf("Starting gRPC gateway server on port " + reverseProxyPort)
	if err := http.ListenAndServe(":" + reverseProxyPort, mux); err != nil {
		log.Fatalf("Could not setup HTTP endpoint: %v", err)
	}
}
