package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//this User model governs all objects inserted or retrieved in the mongo database

// type User struct {
// 	ID            primitive.ObjectID `bson:"_id"`
// 	Email         *string            `json:"email" validate:"required"`
// 	Password      *string            `json:"password" validate:"required,min=6"`
// 	Token         *string            `json:"token"`
// 	Refresh_token *string            `json:"refresh_token"`
// 	Created_at    time.Time          `json:"created_at"`
// 	Updated_at    time.Time          `json:"updated_at"`
// 	User_id       string             `json:"user_id"`
// }

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Username      *string            `json:"username" validate:"required,min=2,max=100"`
	Password      *string            `json:"Password" validate:"required,min=6"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
