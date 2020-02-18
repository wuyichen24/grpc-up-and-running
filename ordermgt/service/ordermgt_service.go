package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	pb "ordergmt/service/ecommerce"
	"strings"
	"time"
)

const (
	maxBatchSize = 3
)

var orderMap = make(map[string]pb.Order)

type orderMgtServer struct {
	orderMap map[string]*pb.Order
}

// Add a new order.
// Simple RPC
func (s *orderMgtServer) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrapper.StringValue, error) {
	if orderReq.Id == "-1" {
		log.Printf("Order ID is invalid! -> Received Order ID %s", orderReq.Id)

		errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
		ds, err := errorStatus.WithDetails(
			&epb.BadRequest_FieldViolation{
				Field:"ID",
				Description: fmt.Sprintf("Order ID received is not valid %s : %s", orderReq.Id, orderReq.Description),
			},
		)
		if err != nil {
			return nil, errorStatus.Err()
		}

		return nil, ds.Err()
	} else {
		log.Printf("Order Added. ID : %v", orderReq.Id)
		orderMap[orderReq.Id] = *orderReq
		return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
	}
}

// Get a order by order ID.
// Simple RPC
func (s *orderMgtServer) GetOrder(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	ord, exists := orderMap[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", orderId)
}

// Get all the orders which has a certain item.
// All the matched orders will be returned from orderMgtServer as a stream.
// Server-side Streaming RPC
func (s *orderMgtServer) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}

// Update multiple orders.
// All the orders will be sent from client as a stream.
// Client-side Streaming RPC
func (s *orderMgtServer) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream.
			return stream.SendAndClose(&wrapper.StringValue{Value: "Orders processed " + ordersStr})
		}

		if err != nil {
			return err
		}
		// Update order
		orderMap[order.Id] = *order

		log.Printf("Order ID : %s - %s", order.Id, "Updated")
		ordersStr += order.Id + ", "
	}
}

// Process multiple orders
// All the order IDs will be sent from client as a stream.
// A combined shipment will contains all the orders which will be delivered to the same destination.
// When the max batch size is reached, all the currently created combined shipments will be sent back to the client.
// Bi-directional Streaming RPC
func (s *orderMgtServer) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	currentBatchSize := 1
	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order : %s", orderId)
		if err == io.EOF {
			// If the stream reached the end (EOF is the signal for the end of the stream)
			// Return all the remaining combined shipments to client.
			log.Printf("EOF : %s", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}

		destination := orderMap[orderId.GetValue()].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			// If the combined shipment has been found for that order by the same destination,
			// Append the order into the combined shipment.
			ord := orderMap[orderId.GetValue()]
			shipment.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			// If the combined shipment hasn't been found for that order by the same destination,
			// Create a new combined shipment, append the order into it.
			comShip := pb.CombinedShipment{Id: "cmb - " + (orderMap[orderId.GetValue()].Destination), Status: "Processed!", }
			ord := orderMap[orderId.GetValue()]
			comShip.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrdersList), comShip.GetId())
		}

		if currentBatchSize == maxBatchSize {
			// If the current batch size reaches the max batch size,
			// return all the combined shipments to client.
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v" , comb.Id, len(comb.OrdersList))
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			currentBatchSize = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			currentBatchSize++
		}
	}
}

// Unary Interceptor (orderMgtServer-side)
// This interceptor consists of pre-processing logic which will be executed before running the remote method,
// post-processing logic which will be executed after running the remote method.
func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Pre-processing phase
	// Gets info about the current RPC call by examining the args passed in
	log.Println("======= [Server Interceptor] ", info.FullMethod)
	log.Printf(" Pre Proc Message : %s", req)

	// Invoking the handler to complete the normal execution of a unary RPC.
	m, err := handler(ctx, req)

	// Post-processing phase
	log.Printf(" Post Proc Message : %s", m)
	return m, err
}

// wrappedStream wraps grpc.ServerStream and intercepts the RecvMsg and SendMsg method call.
type wrappedStream struct {
	grpc.ServerStream
}

// Implementing the RecvMsg function of the wrapper to process messages received with stream RPC.
func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("====== [Server Stream Interceptor Wrapper] Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

// Implementing the SendMsg function of the wrapper to process messages sent with stream RPC.
func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("====== [Server Stream Interceptor Wrapper] Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

// Creating an instance of the new wrapper stream.
func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// Streaming interceptor implementation.
func orderServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Pre-processing phase
	log.Println("====== [Server Stream Interceptor] ", info.FullMethod)

	// Invoking the StreamHandler to complete the execution of RPC invocation
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}