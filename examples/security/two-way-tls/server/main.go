package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-up-and-running/examples/security/two-way-tls/server/ecommerce"
	"io/ioutil"
	"log"
	"net"
)

type server struct {}

var (
	port = ":50051"
	crtFile = "server.crt"    // server public certificate.
	keyFile = "server.key"    // server private key.
	caFile = "ca.crt"         // public certificate of a CA used to sign all public certificates.
)

func main() {
	cert, err := tls.LoadX509KeyPair(crtFile,keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certificate")
	}

	opts := []grpc.ServerOption{
		grpc.Creds(
			credentials.NewTLS(&tls.Config {
				ClientAuth:   tls.RequireAndVerifyClientCert,
				Certificates: []tls.Certificate{cert},
				ClientCAs:    certPool,
			},
		)),
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