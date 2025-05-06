package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product represents a product in the inventory.
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	CategoryID  primitive.ObjectID `bson:"category_id"`
	Stock       int                `bson:"stock"`
}

// Category represents a category of products.
type Category struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

// Inventory represents the structure for inventory management.
type Inventory struct {
	Products  []Product  `bson:"products"`
	Categories []Category `bson:"categories"`
}