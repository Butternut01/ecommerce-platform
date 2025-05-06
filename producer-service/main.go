package main

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	natsURL := os.Getenv("NATS_URL")
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	log.Println("Producer Service is running and connected to NATS")
	select {} // Keep the service running
}
