package model

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Blog ...
@desc representation of blogs schema on database
*/
type Blog struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	UserID    primitive.ObjectID `bson:"userID" json:"userID"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
	UpdatedAt string             `bson:"updatedAt" json:"updatedAt"`
}

/*FromJSON ...
@desc decode request json to Blog struct
*/
func (blog *Blog) FromJSON(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(blog)

	return err
}

/*ToJSON ...
@desc encode Blog struct to JSON
*/
func (blog *Blog) ToJSON() string {
	return "hello"
}
