package main

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	pb "grpc-up-and-running/examples/security/oauth2/server/ecommerce"
	"log"
	"net"
	"strings"
)

type server struct {}

var (
	port               = ":50051"
	crtFile            = "server.crt"    // server public certificate.
	keyFile            = "server.key"    // server private key.

	correctToken       = "some-secret-token"

	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid credentials")
)

func main() {
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	opts := []grpc.ServerOption{
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(ensureValidToken),
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

// This method ensures a valid token exists within a request's metadata.
// - If the token is missing or invalid, the interceptor blocks execution of the handler and returns an error.
// - Otherwise, the interceptor invokes the unary handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	return handler(ctx, req)
}

// Validates the token.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	return token == correctToken
}

func (s server) AddProduct(context.Context, *pb.Product) (*pb.ProductID, error) {
	panic("implement me")
}

func (s server) GetProduct(context.Context, *pb.ProductID) (*pb.Product, error) {
	panic("implement me")
}