package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Relations ...
@desc representation of relations schema on database
*/
type Relations struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	UserID string             `bson:"userID" json:"userID"`
	// List of Followed users ID
	FollowedList string `bson:"followedList" json:"followedList"`
	// List of Blocked users ID
	BlockedList string `bson:"blockedList" json:"blockedList"`
	CreatedAt   string `bson:"createdAt" json:"createdAt"`
	UpdatedAt   string `bson:"updatedAt" json:"updatedAt"`
}
