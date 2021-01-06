package model

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Relations ...
@desc representation of relations schema on database
*/
type Relations struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	UserID primitive.ObjectID `bson:"userID" json:"userID"`
	// List of Followed users ID
	FollowedList []primitive.ObjectID `bson:"followedList" json:"followedList"`
	// List of Blocked users ID
	BlockedList []primitive.ObjectID `bson:"blockedList" json:"blockedList"`
	CreatedAt   string               `bson:"createdAt" json:"createdAt"`
	UpdatedAt   string               `bson:"updatedAt" json:"updatedAt"`
}

/*FromJSON ...
@desc decode request json to Blog struct
*/
func (relations *Relations) FromJSON(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(relations)

	return err
}

/*ToJSON ...
@desc encode Blog struct to JSON
*/
func (relations *Relations) ToJSON() string {
	return "hello"
}
