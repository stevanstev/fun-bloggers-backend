package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Token ...
@desc representation of tokens schema on database
*/
type Token struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	UserID    primitive.ObjectID `bson:"userID" json:"userID"`
	Token     string             `bson:"token" json:"token"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
	UpdatedAt string             `bson:"updatedAt" json:"updatedAt"`
}
