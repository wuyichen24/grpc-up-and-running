package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "productinfo/client/ecommerce"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a new product
	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode.`
	price := float32(1000.0)

	// Add a new product
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	// Get a product
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: ", product.String())

}
