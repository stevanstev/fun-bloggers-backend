package driver

import (
	"context"

	models "fun-blogger-backend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("funBloggers"), nil
}

/*Insert ...
@desc insert data to it's model
@param collectionName database's collection name
@param insertData data model
*/
func Insert(collectionName string, insertData interface{}) error {
	db, err := connect()
	if err != nil {
		return err
	}

	_, err = db.Collection(collectionName).InsertOne(ctx, insertData)
	if err != nil {
		return err
	}

	return nil
}

/*FindUsers ...
@param query bson.M{}
*/
func FindUsers(query map[string]interface{}) ([]models.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	cursor, err := db.Collection("users").Find(ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	result := make([]models.User, 0)
	for cursor.Next(ctx) {
		var row models.User
		err := cursor.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}

/*FindBlogs ...
@param query bson.M{}
*/
func FindBlogs(query map[string]interface{}) ([]models.Blog, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	cursor, err := db.Collection("blogs").Find(ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	result := make([]models.Blog, 0)
	for cursor.Next(ctx) {
		var row models.Blog
		err := cursor.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}

/*FindTokens ...
@param query bson.M{}
*/
func FindTokens(query map[string]interface{}) ([]models.Token, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	cursor, err := db.Collection("tokens").Find(ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	result := make([]models.Token, 0)
	for cursor.Next(ctx) {
		var row models.Token
		err := cursor.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}

/*GetUserIDByToken ...
@desc find user id by token value
@param token, token header
*/
func GetUserIDByToken(token string) primitive.ObjectID {
	query := bson.M{"token": token}

	result, _ := FindTokens(query)

	return result[0].UserID
}
