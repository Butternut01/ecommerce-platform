package main

import (
    "log"
    "time"

    "github.com/gin-gonic/gin"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "github.com/Butternut01/ecommerce-platform/api-gateway/handlers"
    "github.com/Butternut01/ecommerce-platform/api-gateway/proto/inventory"
    "github.com/Butternut01/ecommerce-platform/api-gateway/proto/order"
    "github.com/Butternut01/ecommerce-platform/api-gateway/proto/user"
)

func main() {
    // Set Gin to release mode
    gin.SetMode(gin.ReleaseMode)

    // Initialize gin router
    r := gin.Default()

    // Set up gRPC connection options
    opts := []grpc.DialOption{
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithBlock(),
        grpc.WithTimeout(30 * time.Second), // Increased timeout
    }

    // Initialize gRPC clients with Docker service names
    inventoryConn, err := grpc.Dial("inventory-service:50051", opts...)
    if err != nil {
        log.Fatalf("Failed to connect to Inventory service: %v", err)
    }
    defer inventoryConn.Close()
    inventoryClient := inventory.NewInventoryServiceClient(inventoryConn)

    orderConn, err := grpc.Dial("order-service:50052", opts...)
    if err != nil {
        log.Fatalf("Failed to connect to Order service: %v", err)
    }
    defer orderConn.Close()
    orderClient := order.NewOrderServiceClient(orderConn)

    userConn, err := grpc.Dial("user-service:50053", opts...)
    if err != nil {
        log.Fatalf("Failed to connect to User service: %v", err)
    }
    defer userConn.Close()
    userClient := user.NewUserServiceClient(userConn)

    // Initialize the gateway handler
    gatewayHandler := handlers.NewGatewayHandler(inventoryClient, orderClient, userClient)

    // Define API routes
    r.POST("/products", gatewayHandler.CreateProduct)
    r.GET("/products/:id", gatewayHandler.GetProduct)
    r.POST("/orders", gatewayHandler.CreateOrder)
    r.GET("/orders/:id", gatewayHandler.GetOrder)
    r.POST("/users/register", gatewayHandler.RegisterUser)
    r.POST("/users/login", gatewayHandler.LoginUser)

    // Start the server
    log.Println("API Gateway starting on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}