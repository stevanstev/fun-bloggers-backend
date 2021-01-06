package model

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User ...
@desc representation of users schema on database
*/
type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	FullName  string             `bson:"fullName" json:"fullName"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
	UpdatedAt string             `bson:"updatedAt" json:"updatedAt"`
}

/*FromJSON ...
@desc decode request json to User struct
*/
func (user *User) FromJSON(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)

	return err
}

/*ToJSON ...
@desc encode User struct to JSON
*/
func (user *User) ToJSON() string {
	return "json"
}
