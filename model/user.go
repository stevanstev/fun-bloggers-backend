package model

import (
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	BaseModel
	Id primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	FullName string `bson:"fullName" json:"fullName"`
	CreatedAt string `bson:"createdAt" json: "createdAt"`
	UpdatedAt string `bson:"updatedAt" json: "updatedAt"`
}

func (user *User) FromJson(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)

	return err
}

func (user *User) ToJson() string {
	return "json"
}