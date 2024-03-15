package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Image string             `bson:"image" json:"image"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}
