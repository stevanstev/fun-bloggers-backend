package model

import (
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	BaseModel
	Id primitive.ObjectID `bson:"_id" json: "_id"`
	AuthorID int `bson:"authorId" json: "authorId"`
	Title string `bson:"title" json: "title"`
	Content string `bson:"content" json: "content"`
	CreatedAt string `bson:"createdAt" json: "createdAt"`
	UpdatedAt string `bson:"updatedAt" json: "updatedAt"`
}

func (blog *Blog) FromJSON(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(blog)

	return err
}

func (blog *Blog) ToJSON() string {
	return "hello"
}