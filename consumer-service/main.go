package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	pb "github.com/Butternut01/ecommerce-platform/inventory-service/proto"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	// Connect to NATS
	natsURL := os.Getenv("NATS_URL")
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Connect to Inventory Service
	inventoryAddr := os.Getenv("INVENTORY_SERVICE_ADDR")
	conn, err := grpc.Dial(inventoryAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Inventory Service: %v", err)
	}
	defer conn.Close()
	inventoryClient := pb.NewInventoryServiceClient(conn)

	// Subscribe to order.created events
	_, err = nc.Subscribe("order.created", func(msg *nats.Msg) {
		var event map[string]interface{}
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Failed to parse event: %v", err)
			return
		}

		// Call Inventory Service to decrease stock
		for _, item := range event["items"].([]interface{}) {
			product := item.(map[string]interface{})
			_, err := inventoryClient.UpdateProduct(context.Background(), &pb.UpdateProductRequest{
				Id:       product["product_id"].(string),
				Quantity: int32(-product["quantity"].(float64)),
			})
			if err != nil {
				log.Printf("Failed to update product stock: %v", err)
			}
		}

		log.Printf("Processed order.created event: %v", event)
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to order.created: %v", err)
	}

	log.Println("Consumer Service is running and subscribed to order.created")
	select {} // Keep the service running
}
