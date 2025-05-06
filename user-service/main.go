package main

import (
	"log"
	"net"

	"github.com/Butternut01/ecommerce-platform/user-service/handlers"
	pb "github.com/Butternut01/ecommerce-platform/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Set up a listener on TCP port 50053
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the User service with the gRPC server
	pb.RegisterUserServiceServer(s, &handlers.UserService{})

	// Enable gRPC reflection
	reflection.Register(s)

	log.Println("User Service is running on port :50053")

	// Start the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
