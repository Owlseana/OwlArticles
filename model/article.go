package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Author    string             `bson:"author" json:"author" validate:"required"`
	Title     string             `bson:"title" json:"title" validate:"required"`
	Content   string             `bson:"content" json:"content" validate:"required"`
	CreatedAt string             `bson:"created_at" json:"created_at"`
	UpdatedAt string             `bson:"updated_at" json:"updated_at"`
}
