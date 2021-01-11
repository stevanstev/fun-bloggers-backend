package driver

import (
	"context"
	"log"

	models "fun-blogger-backend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"fmt"
)

var ctx = context.Background()

type blogResultType struct {
	ID        primitive.ObjectID `json:"_id"`
	UserID    primitive.ObjectID `json:"userID"`
	Author    string             `json:"author"`
	CreatedAt string             `json:"createdAt"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()

	password := "<PASSWORD>"
	username := "<USERNAME>"
	dbName := "<DBNAME>"
	uri := fmt.Sprintf("mongodb+srv://%s:%s@mycluster.waynb.mongodb.net/%s?retryWrites=true&w=majority", username, password, dbName)

	clientOptions.ApplyURI(uri)
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
func FindBlogs(query map[string]interface{}) ([]blogResultType, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	cursor, err := db.Collection("blogs").Find(ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	result := make([]blogResultType, 0)
	for cursor.Next(ctx) {
		var row models.Blog
		err := cursor.Decode(&row)
		if err != nil {
			return nil, err
		}

		query := bson.M{
			"_id": row.UserID,
		}

		user, _ := FindUsers(query)

		var blogResult blogResultType

		blogResult.ID = row.ID
		blogResult.UserID = row.UserID
		blogResult.Author = user[0].Email
		blogResult.CreatedAt = row.CreatedAt
		blogResult.Title = row.Title
		blogResult.Content = row.Content
		blogResult.Content = row.Content

		result = append(result, blogResult)
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

/*FindRelations ...
@param query bson.M{}
*/
func FindRelations(query map[string]interface{}) ([]models.Relations, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	cursor, err := db.Collection("relations").Find(ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	result := make([]models.Relations, 0)
	for cursor.Next(ctx) {
		var row models.Relations
		err := cursor.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}

/*DeleteToken ...
 */
func DeleteToken(token string) error {
	db, err := connect()
	if err != nil {
		return err
	}

	query := bson.M{
		"token": token,
	}

	_, err = db.Collection("tokens").DeleteOne(ctx, query)

	if err != nil {
		return err
	}

	return nil
}

/*GetUserIDByToken ...
@desc find user id by token value
@param token, token header
*/
func GetUserIDByToken(token string) primitive.ObjectID {
	query := bson.M{"token": token}

	result, _ := FindTokens(query)

	if len(result) == 0 {
		return primitive.NilObjectID
	}

	return result[0].UserID
}

/*UpdateRelations ...
@desc update user by token query
@param token, token header
*/
func UpdateRelations(collection string, where map[string]interface{}, updates map[string]interface{}) error {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection(collection).UpdateOne(ctx, where, updates)

	if err != nil {
		return err
	}

	return nil
}

/*GetListOfBlockedUsers ...
@desc get list of blocked users
@param query, find list of blocked users by query specified
*/
func GetListOfBlockedUsers(query map[string]interface{}) ([]models.User, error) {
	var blockedUserResult []models.User

	relationsResult, _ := FindRelations(query)

	blockedUsers := relationsResult[0].BlockedList

	if len(blockedUsers) == 0 {
		return []models.User{}, nil
	}

	for i := 0; i < len(blockedUsers); i++ {
		query := bson.M{
			"userID": blockedUsers[i],
		}
		user, _ := FindUsers(query)
		blockedUserResult = append(blockedUserResult, user[0])
	}

	return blockedUserResult, nil
}

/*GetListOfFollowedUsers ...
@desc get list of followed users
@param query, find list of followed users by query specified
*/
func GetListOfFollowedUsers(query map[string]interface{}) ([]models.User, error) {
	var followedUserResult []models.User

	relationsResult, err := FindRelations(query)

	if err != nil {
		return nil, err
	}

	followedUsers := relationsResult[0].FollowedList

	if len(followedUsers) == 0 {
		return []models.User{}, nil
	}

	for i := 0; i < len(followedUsers); i++ {
		query := bson.M{
			"_id": followedUsers[i],
		}
		user, _ := FindUsers(query)
		followedUserResult = append(followedUserResult, user[0])
	}

	return followedUserResult, nil
}
