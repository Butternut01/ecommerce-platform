package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/Butternut01/ecommerce-platform/order-service/proto"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderHandler struct {
	orderClient pb.OrderServiceClient
}

// OrderService struct
type OrderService struct {
	client *mongo.Client
	nc     *nats.Conn
	pb.UnimplementedOrderServiceServer
}

func NewOrderHandler(conn *grpc.ClientConn) *OrderHandler {
	return &OrderHandler{
		orderClient: pb.NewOrderServiceClient(conn),
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating an order
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for retrieving an order
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an order
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an order
}

// NewOrderService creates a new OrderService
func NewOrderService(client *mongo.Client, nc *nats.Conn) *OrderService {
	return &OrderService{client: client, nc: nc}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Define and initialize newOrder
	newOrder := struct {
		ID     struct{ Hex func() string }
		UserID struct{ Hex func() string }
		Items  []interface{}
	}{
		ID:     struct{ Hex func() string }{Hex: func() string { return "order123" }},
		UserID: struct{ Hex func() string }{Hex: func() string { return "user456" }},
		Items:  []interface{}{"item1", "item2"},
	}

	// Publish order.created event to NATS
	event := map[string]interface{}{
		"order_id": newOrder.ID.Hex(),
		"user_id":  newOrder.UserID.Hex(),
		"items":    newOrder.Items,
	}
	eventData, _ := json.Marshal(event)
	if err := s.nc.Publish("order.created", eventData); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to publish event: %v", err)
	}

	return &pb.CreateOrderResponse{
		OrderId: newOrder.ID.Hex(),
		Status:  "created",
	}, nil
}
