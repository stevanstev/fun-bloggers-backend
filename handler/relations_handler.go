package handler

import (
	"encoding/json"
	driver "fun-blogger-backend/driver"
	library "fun-blogger-backend/library"
	models "fun-blogger-backend/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*RelationsHandlerGet ...
@desc handling get request of /relations
@route /relations
@method GET
@access Private
*/
func RelationsHandlerGet(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)

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
	library.SetDefaultHTTPHeader(w)

	reqToken := r.Header.Get("x-auth-token")
	userID := driver.GetUserIDByToken(reqToken)

	var responsesMap = make(map[string]interface{}, 0)

	query := bson.M{
		"userID": userID,
	}

	followedUsers, err := driver.GetListOfFollowedUsers(query)

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if len(followedUsers) == 0 {
		responsesMap["followedUsers"] = "There are no users that you follow"
	} else {
		responsesMap["followedUsers"] = followedUsers
	}

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
	library.SetDefaultHTTPHeader(w)

	reqToken := r.Header.Get("x-auth-token")
	userID := driver.GetUserIDByToken(reqToken)

	var responsesMap = make(map[string]interface{}, 0)

	query := bson.M{
		"userID": userID,
	}

	blockedUsers, err := driver.GetListOfBlockedUsers(query)

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if len(blockedUsers) == 0 {
		responsesMap["blockedUsers"] = "There are no users that you block"
	} else {
		responsesMap["blockedUsers"] = blockedUsers
	}

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}

/*RelationsByEmailHandlerPost ...
@desc handling post request of /relations/followers
@route /relations/followers
@method POST
@access Private
*/
func RelationsByEmailHandlerPost(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)

	var responsesMap = make(map[string]interface{}, 0)
	var user models.User
	err := user.FromJSON(r)

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if user.Email == "" {
		responsesMap["email"] = "Email cannot be Empty"
		responsesMap["status"] = "false"
	} else {
		query := bson.M{
			"email": user.Email,
		}

		findUserByEmail, _ := driver.FindUsers(query)

		userID := findUserByEmail[0].ID

		query = bson.M{
			"userID": userID,
		}

		followedUsers, err := driver.GetListOfFollowedUsers(query)

		if err != nil {
			library.ResponseByCode(500, w, err.Error())
			return
		}

		responsesMap["count"] = len(followedUsers)

		if responsesMap["count"] == 0 {
			responsesMap["followers"] = "This user does not follow anyone"
		} else {
			responsesMap["followers"] = followedUsers
		}

		responsesMap["success"] = "true"
	}

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}

//API FOR BLOCKING USER AND FOLLOWING USER
// Sam won't be able to see any of kim's blog post.
// Sam is not allowed to follow Kim anymore. Note: User Story (6)

/*RelationsBlockingUserHandlerPost ...
@desc handling post request of /relations/block
@route /relations/block
@method POST
@access Private
*/
func RelationsBlockingUserHandlerPost(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)
	type usersEmail struct {
		UserEmail     string `json:"userEmail"`
		FollowerEmail string `json:"followerEmail"`
	}

	var emails usersEmail
	var responsesMap = make(map[string]interface{}, 0)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&emails)

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if emails.UserEmail == "" {
		responsesMap["userEmail"] = "User email cannot be empty"
	}

	if emails.FollowerEmail == "" {
		responsesMap["followerEmail"] = "Follower email cannot be empty"
	}

	if len(responsesMap) == 0 {
		// Find data and block
		query := bson.M{
			"$or": []bson.M{
				bson.M{
					"email": emails.UserEmail,
				},
				bson.M{
					"email": emails.FollowerEmail,
				},
			},
		}

		result, _ := driver.FindUsers(query)

		switch len(result) {
		case 0:
			responsesMap["emailErorr"] = "Email not found"
		case 1:
			responsesMap["emailErorr"] = "Either userEmail or followerEmail not found"
		default:
			var userDetail models.User
			var followerDetail models.User

			for i := 0; i < len(result); i++ {
				if result[i].Email == emails.UserEmail {
					userDetail = result[i]
				} else {
					followerDetail = result[i]
				}
			}

			// Add new blocked user to sender's list
			query := bson.M{
				"userID": userDetail.ID,
			}
			userRelation, _ := driver.FindRelations(query)
			blockedList := userRelation[0].BlockedList
			blockedList = append(blockedList, followerDetail.ID)

			query = bson.M{"$set": bson.M{"blockedList": blockedList}}

			where := bson.M{"userID": userDetail.ID}

			err = driver.UpdateRelations("relations", where, query)

			if err != nil {
				library.ResponseByCode(500, w, err.Error())
				return
			}

			// Remove followed user of target's account (if target already followed the sender)

			responsesMap["status"] = "true"
		}
	} else {
		responsesMap["status"] = "false"
	}

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}

/*RelationsFollowingUserHandlerPost ...
@desc handling post request of /relations/follow
@route /relations/follow
@method POST
@access Private
*/
func RelationsFollowingUserHandlerPost(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)

	var responsesMap = make(map[string]interface{}, 0)

	var userToFollow models.User
	userToFollow.FromJSON(r)

	if userToFollow.Email == "" {
		responsesMap["emailError"] = "Email cannot be empty"
	}

	if len(responsesMap) != 0 {
		responsesMap["status"] = "false"
	} else {
		query := bson.M{
			"email": userToFollow.Email,
		}

		followingUserDetails, _ := driver.FindUsers(query)

		// Get sender relation
		reqToken := r.Header.Get("x-auth-token")
		userID := driver.GetUserIDByToken(reqToken)

		query = bson.M{
			"userID": userID,
		}

		userRelation, _ := driver.FindRelations(query)
		followedList := userRelation[0].FollowedList
		followedList = append(followedList, followingUserDetails[0].ID)

		query = bson.M{"$set": bson.M{"followedList": followedList}}

		where := bson.M{"userID": userID}

		err := driver.UpdateRelations("relations", where, query)

		if err != nil {
			library.ResponseByCode(500, w, err.Error())
			return
		}

		responsesMap["status"] = "true"
	}

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}

/*FindCommonFollowersHandlerPost ...
@desc handling post request of /relations/common
@route /relations/common
@method POST
@access Private
*/
func FindCommonFollowersHandlerPost(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)
	var responsesMap = make(map[string]interface{}, 0)

	type listOfEmails struct {
		Emails []string `json:"emails"`
	}

	var usersEmail listOfEmails
	var userIDS []primitive.ObjectID

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&usersEmail)

	if len(usersEmail.Emails) == 0 {
		responsesMap["emailError"] = "Emails cannot be empty"
	} else if len(usersEmail.Emails) < 2 {
		responsesMap["emailError"] = "Please specify at least 2 emails"
	}

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if len(responsesMap) != 0 {
		responsesMap["success"] = "false"
		responsesMap["count"] = 0
	} else {
		for _, e := range usersEmail.Emails {
			query := bson.M{
				"email": e,
			}

			user, _ := driver.FindUsers(query)

			userIDS = append(userIDS, user[0].ID)
		}

		type usersFollowedList struct {
			FollowedList [][]primitive.ObjectID
		}

		var listOfFollowedList usersFollowedList

		for _, i := range userIDS {
			query := bson.M{
				"userID": i,
			}
			userRelation, _ := driver.FindRelations(query)

			listOfFollowedList.FollowedList = append(listOfFollowedList.FollowedList, userRelation[0].FollowedList)
		}

		searchRepetition := len(listOfFollowedList.FollowedList)
		var initCommonID []primitive.ObjectID
		initCommonID = listOfFollowedList.FollowedList[0]

		var searchCommon []primitive.ObjectID
		canContinue := true

		if len(initCommonID) == 0 {
			responsesMap["followers"] = "No common followers"
		} else {
			for i := 0; i < len(initCommonID); i++ {
				for j := 1; j < searchRepetition; j++ {
					if canContinue == true {
						if len(listOfFollowedList.FollowedList[j]) == 0 {
							canContinue = false
						} else {
							for k := 0; k < len(listOfFollowedList.FollowedList[j]); k++ {
								if initCommonID[i] == listOfFollowedList.FollowedList[j][k] {
									canContinue = true
									break
								}

								canContinue = false
							}
						}
					} else {
						break
					}

					if j == searchRepetition-1 && canContinue != false {
						searchCommon = append(searchCommon, initCommonID[i])
					}
				}
				canContinue = true
			}

			var usersEmailCommon []string
			if len(searchCommon) == 0 {
				responsesMap["followers"] = "No common followers"
			} else {
				for _, id := range searchCommon {
					query := bson.M{
						"_id": id,
					}
					getUser, _ := driver.FindUsers(query)

					usersEmailCommon = append(usersEmailCommon, getUser[0].Email)
				}

				responsesMap["followers"] = usersEmailCommon
			}

			responsesMap["count"] = len(searchCommon)
			responsesMap["success"] = "true"
		}
	}

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}
