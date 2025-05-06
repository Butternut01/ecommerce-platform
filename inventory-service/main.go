package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/Butternut01/ecommerce-platform/inventory-service/handlers"
	pb "github.com/Butternut01/ecommerce-platform/inventory-service/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Configure logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Inventory Service...")

	// Set up MongoDB connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Verify MongoDB connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Successfully connected to MongoDB")

	// Set up gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server with health check
	grpcServer := grpc.NewServer()

	// Register health service
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	log.Println("Health check service registered and set to SERVING")

	// Register inventory service
	inventoryService := handlers.NewInventoryService(client)
	pb.RegisterInventoryServiceServer(grpcServer, inventoryService)

	// Enable gRPC reflection
	reflection.Register(grpcServer)

	log.Println("Inventory Service is running on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
