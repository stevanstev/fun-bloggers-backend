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
	library.SetDefaultHTTPHeader(&w)

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
	library.SetDefaultHTTPHeader(&w)

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
	library.SetDefaultHTTPHeader(&w)

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

/*RelationsHandlerUserGet ...
@desc handling post request of /relations/user
@route /relations/user
@method GET
@access Private
*/
func RelationsHandlerUserGet(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(&w)

	reqToken := r.Header.Get("x-auth-token")
	userID := driver.GetUserIDByToken(reqToken)

	var responsesMap = make(map[string]interface{}, 0)

	query := bson.M{
		"userID": userID,
	}

	relations, _ := driver.FindRelations(query)

	if len(relations) != 0 {
		followedUser := len(relations[0].FollowedList)
		blockedUser := len(relations[0].BlockedList)

		responsesMap["followedUser"] = followedUser
		responsesMap["blockedUser"] = blockedUser

		if followedUser == 0 {
			responsesMap["followedUser"] = 0
		}

		if blockedUser == 0 {
			responsesMap["blockedUser"] = 0
		}
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
	library.SetDefaultHTTPHeader(&w)

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

		findUserByEmail, err := driver.FindUsers(query)

		if err != nil {
			library.ResponseByCode(500, w, err.Error())
			return
		}

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
			responsesMap["following"] = "User does not follow anyone"
		} else {
			responsesMap["following"] = followedUsers
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
	library.SetDefaultHTTPHeader(&w)
	type usersEmail struct {
		UserEmail  string `json:"userEmail"`
		BlockEmail string `json:"blockEmail"`
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

	if emails.BlockEmail == "" {
		responsesMap["blockEmail"] = "Block email cannot be empty"
	}

	if len(responsesMap) == 0 {
		// Find data and block
		query := bson.M{
			"$or": []bson.M{
				bson.M{
					"email": emails.UserEmail,
				},
				bson.M{
					"email": emails.BlockEmail,
				},
			},
		}

		result, _ := driver.FindUsers(query)

		switch len(result) {
		case 0:
			responsesMap["emailErorr"] = "Email not found"
		case 1:
			responsesMap["emailErorr"] = "Either userEmail or blockEmail not found"
		default:
			var userDetails models.User
			var blockUserDetails models.User

			for i := 0; i < len(result); i++ {
				if result[i].Email == emails.UserEmail {
					userDetails = result[i]
				} else {
					blockUserDetails = result[i]
				}
			}

			// Add new blocked user to sender's list
			query := bson.M{
				"userID": userDetails.ID,
			}
			userRelation, _ := driver.FindRelations(query)
			blockedList := userRelation[0].BlockedList
			blockedList = append(blockedList, blockUserDetails.ID)

			//Remove blocked user's id from sender followedlist
			followedList := userRelation[0].FollowedList
			var updatedFollowedList []primitive.ObjectID

			for i := 0; i < len(followedList); i++ {
				if followedList[i] == blockUserDetails.ID {
					continue
				}

				updatedFollowedList = append(updatedFollowedList, followedList[i])
			}

			query = bson.M{"$set": bson.M{"blockedList": blockedList, "followedList": updatedFollowedList}}

			where := bson.M{"userID": userDetails.ID}

			err = driver.UpdateRelations("relations", where, query)

			if err != nil {
				library.ResponseByCode(500, w, err.Error())
				return
			}

			// Remove user in blocked user's followed list where  user id equals to sender user id
			blockedUserID := blockUserDetails.ID

			query = bson.M{
				"userID": blockedUserID,
			}

			blockedUserRelations, _ := driver.FindRelations(query)

			blockedUserFollowedList := blockedUserRelations[0].FollowedList
			blockedUsrFolListLen := len(blockedUserFollowedList)

			var updateBlckdUsrFllwdList []primitive.ObjectID

			if blockedUsrFolListLen != 0 {
				for i := 0; i < blockedUsrFolListLen; i++ {
					if blockedUserFollowedList[i] == userDetails.ID {
						continue
					}

					updateBlckdUsrFllwdList = append(updateBlckdUsrFllwdList, blockedUserFollowedList[i])
				}
			}

			query = bson.M{"$set": bson.M{"followedList": updateBlckdUsrFllwdList}}

			where = bson.M{"userID": blockedUserID}

			err = driver.UpdateRelations("relations", where, query)

			if err != nil {
				library.ResponseByCode(500, w, err.Error())
				return
			}

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
	library.SetDefaultHTTPHeader(&w)

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

/*CheckIfUserAlreadyFollowedPost ...
@desc handling post request of /relations/already-following
@route /relations/already-following
@method POST
@access Private
*/
func CheckIfUserAlreadyFollowedPost(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(&w)

	reqToken := r.Header.Get("x-auth-token")
	var responsesMap = make(map[string]interface{}, 0)
	userID := driver.GetUserIDByToken(reqToken)

	query := bson.M{
		"userID": userID,
	}

	userRelation, _ := driver.FindRelations(query)

	var targetUser models.User

	targetUser.FromJSON(r)

	if targetUser.ID == primitive.NilObjectID {
		responsesMap["idError"] = "User ID is empty"
	}

	if len(responsesMap) != 0 {
		responsesMap["status"] = "false"
	} else {
		isFollowed := "no"

		if targetUser.ID == userRelation[0].UserID {
			//Follow itself
			isFollowed = "abort"
		} else {
			//do query
			followedList := userRelation[0].FollowedList
			for i := 0; i < len(followedList); i++ {
				if targetUser.ID == followedList[i] {
					isFollowed = "yes"
					break
				}
			}
		}

		responsesMap["status"] = isFollowed
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
	library.SetDefaultHTTPHeader(&w)
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
