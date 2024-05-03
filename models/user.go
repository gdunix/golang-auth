package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	ID      primitive.ObjectID `bson:"_id"`
	Username *string            `json:"username" validate:"required,min=6"`
	Email    *string            `json:"email"`
	Password *string            `json:"password" validate:"required,min=6"`
}
