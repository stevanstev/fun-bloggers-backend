package handler

import (
	"encoding/json"
	"fmt"
	driver "fun-blogger-backend/driver"
	library "fun-blogger-backend/library"
	models "fun-blogger-backend/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)

/*RegisterHandlerGet ...
@desc handling get request of /register
@route /register
@method GET
@access Public
*/
func RegisterHandlerGet(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)
	
	response := models.BaseResponse{}
	response.GetDefault("Register Api Ready")

	httpResponse, err := json.Marshal(response)

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	fmt.Fprint(w, string(httpResponse))
}

/*RegisterHandlerPost ...
@desc handling post request of /register
@route /register
@method POST
@access Public
*/
func RegisterHandlerPost(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)
	
	var user models.User
	err := user.FromJSON(r)
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	var responsesMap = map[string]string{}

	isEmailValid := func(e string) bool {
		if len(e) < 3 && len(e) > 254 {
			return false
		}
		return emailRegex.MatchString(e)
	}


	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if user.Email == "" {
		responsesMap["email"] = "Email cannot be Empty"	
	}

	if user.Email != "" {
		if !isEmailValid(user.Email) {
			responsesMap["email"] = "Please specify correct email format"
		}
	}

	if user.Password == "" {
		responsesMap["password"] = "Password cannot be Empty"
	}

	if user.FullName == "" {
		responsesMap["fullName"] = "Fullname cannot be Empty"
	}

	if len(responsesMap) != 0 {
		responsesMap["status"] = "false"
	} else {
		query := bson.M{"email": user.Email}
		findUser, _ := driver.FindUsers(query)

		if len(findUser) != 0 {
			responsesMap["email"] = "Please use another email address"
			responsesMap["status"] = "false"
		} else {
			// Set User ID
			user.ID = primitive.NewObjectID()
			user.CreatedAt = library.GetCurrentDate()
			user.UpdatedAt = library.GetCurrentDate()
			user.Password, _ = library.EncryptPassword([]byte(user.Password))

			// Insert model to users table
			err = driver.Insert("users", user)

			if err != nil {
				library.ResponseByCode(500, w, "There's something wrong")
				return
			}

			// Create new relation data with empty FollowedList and BlockedList
			var relations models.Relations
			relations.ID = primitive.NewObjectID()
			relations.UserID = user.ID
			relations.FollowedList = []primitive.ObjectID{}
			relations.BlockedList = []primitive.ObjectID{}
			relations.CreatedAt = library.GetCurrentDate()
			relations.UpdatedAt = library.GetCurrentDate()

			err = driver.Insert("relations", relations)

			responsesMap["status"] = "true"
		}
	}

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}
