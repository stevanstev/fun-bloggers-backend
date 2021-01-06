package handler

import (
	"encoding/json"
	driver "fun-blogger-backend/driver"
	library "fun-blogger-backend/library"
	models "fun-blogger-backend/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

/*RelationsHandlerGet ...
@desc handling get request of /relations
@route /relations
@method GET
@access Private
*/
func RelationsHandlerGet(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("x-auth-token")
	userID := driver.GetUserIDByToken(reqToken)

	var responsesMap = make(map[string]interface{}, 0)

	query := bson.M{
		"userID": userID,
	}
	relationsResult, _ := driver.FindRelations(query)

	followedUserCount := len(relationsResult[0].FollowedList)
	blockedUserCount := len(relationsResult[0].BlockedList)

	responsesMap["followedUserCount"] = followedUserCount
	responsesMap["blockedUserCount"] = blockedUserCount

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}

/*RelationsFollowedHandlerGet ...
@desc handling get request of /relations/followed
@route /relations/followed
@method GET
@access Private
*/
func RelationsFollowedHandlerGet(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("x-auth-token")
	userID := driver.GetUserIDByToken(reqToken)

	var responsesMap = make(map[string]interface{}, 0)

	query := bson.M{
		"userID": userID,
	}
	relationsResult, _ := driver.FindRelations(query)

	followedUsers := relationsResult[0].FollowedList

	if len(followedUsers) == 0 {
		responsesMap["followedUsers"] = "There are no users that you follow"
	} else {
		var followedUserResult []models.User

		for i := 0; i < len(followedUsers); i++ {
			query := bson.M{
				"userID": followedUsers[i],
			}
			user, _ := driver.FindUsers(query)
			followedUserResult = append(followedUserResult, user[0])
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}

/*RelationsBlockedHandlerGet ...
@desc handling get request of /relations/blocked
@route /relations/blocked
@method GET
@access Private
*/
func RelationsBlockedHandlerGet(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("x-auth-token")
	userID := driver.GetUserIDByToken(reqToken)

	var responsesMap = make(map[string]interface{}, 0)

	query := bson.M{
		"userID": userID,
	}
	relationsResult, _ := driver.FindRelations(query)

	blockedUsers := relationsResult[0].BlockedList

	if len(blockedUsers) == 0 {
		responsesMap["blockedUsers"] = "There are no users that you block"
	} else {
		var BlockedUserResult []models.User

		for i := 0; i < len(blockedUsers); i++ {
			query := bson.M{
				"userID": blockedUsers[i],
			}
			user, _ := driver.FindUsers(query)
			BlockedUserResult = append(BlockedUserResult, user[0])
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}
