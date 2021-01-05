package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Relations struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	UserId string `bson:"userId" json:"userId"`
	// List of Followed users ID
	FollowedList string `bson:"followedList" json:"followedList"`
	// List of Blocked users ID
	BlockedList string `bson:"blockedList" json:"blockedList"`
	CreatedAt string `bson:"createdAt" json: "createdAt"`
	UpdatedAt string `bson:"updatedAt" json: "updatedAt"`
}