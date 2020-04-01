package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	pb "grpc-up-and-running/examples/security/two-way-tls/client/ecommerce"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
)

var (
	address = "localhost:50051"
	hostname = "localhost"
	crtFile = "client.crt"
	keyFile = "client.key"
	caFile = "ca.crt"
)

func main() {
	cert, err := tls.LoadX509KeyPair(crtFile,keyFile)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certs")
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials( credentials.NewTLS(&tls.Config{
			ServerName:   hostname,                      // ServerName must be equal to the Common Name on the certificate.
			Certificates: []tls.Certificate{cert},
			RootCAs:      certPool,
		})),
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
