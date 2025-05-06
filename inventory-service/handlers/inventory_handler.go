package handlers

import (
	"context"

	"github.com/Butternut01/ecommerce-platform/inventory-service/models"
	pb "github.com/Butternut01/ecommerce-platform/inventory-service/proto"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InventoryHandler struct {
	client pb.InventoryServiceClient
}

// InventoryService represents the gRPC service
type InventoryService struct {
	pb.UnimplementedInventoryServiceServer
	db *mongo.Client
}

func NewInventoryHandler(conn *grpc.ClientConn) *InventoryHandler {
	return &InventoryHandler{
		client: pb.NewInventoryServiceClient(conn),
	}
}

func (h *InventoryHandler) GetInventory(c *gin.Context) {
	// Implementation for getting inventory
}

func (h *InventoryHandler) AddProduct(c *gin.Context) {
	// Implementation for adding a product
}

func (h *InventoryHandler) UpdateProduct(c *gin.Context) {
	// Implementation for updating a product
}

func (h *InventoryHandler) DeleteProduct(c *gin.Context) {
	// Implementation for deleting a product
}

// NewInventoryService creates a new InventoryService
func NewInventoryService(client *mongo.Client) *InventoryService {
	return &InventoryService{db: client}
}

func (s *InventoryService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	// Fetch products from MongoDB
	collection := s.db.Database("ecommerce").Collection("products")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch products: %v", err)
	}
	defer cursor.Close(ctx)

	var products []*pb.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to decode product: %v", err)
		}
		products = append(products, &pb.Product{
			Id:          product.ID.Hex(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    int32(product.Stock),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "Cursor error: %v", err)
	}

	return &pb.ListProductsResponse{Products: products}, nil
}
