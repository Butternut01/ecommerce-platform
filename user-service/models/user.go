package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the e-commerce platform
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

// UserRegistration represents the data required for user registration
type UserRegistration struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserLogin represents the data required for user login
type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}