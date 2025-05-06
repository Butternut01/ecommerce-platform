package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order represents an order in the e-commerce platform.
type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	ProductID   primitive.ObjectID `bson:"product_id"`
	Quantity    int                `bson:"quantity"`
	Status      string             `bson:"status"` // e.g., "pending", "shipped", "delivered"
	CreatedAt   primitive.DateTime `bson:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at"`
}