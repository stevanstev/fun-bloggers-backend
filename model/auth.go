package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	UserId string `bson:"userId", json:"userId"`
	token string  `bson:"token", json:"token"`
	CreatedAt string `bson:"createdAt" json: "createdAt"`
	UpdatedAt string `bson:"updatedAt" json: "updatedAt"`
}