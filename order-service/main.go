package main

import (
	"context"
	"log"
	"net"

	"github.com/Butternut01/ecommerce-platform/order-service/handlers"
	pb "github.com/Butternut01/ecommerce-platform/order-service/proto"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Connect to NATS
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Set up gRPC server
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderService := handlers.NewOrderService(client)
	pb.RegisterOrderServiceServer(grpcServer, orderService)

	// Enable gRPC reflection
	reflection.Register(grpcServer)

	log.Println("Order Service is running on port :50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
