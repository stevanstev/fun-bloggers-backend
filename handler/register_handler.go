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
)

/*RegisterHandlerGet ...
@desc handling get request of /register
@route /register
@method GET
@access Public
*/
func RegisterHandlerGet(w http.ResponseWriter, r *http.Request) {
	response := models.BaseResponse{}
	response.GetDefault("Register Api Ready")

	w.Header().Add("Content-Type", "application/json")
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
	var user models.User
	err := user.FromJSON(r)

	var responsesMap = map[string]string{}

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if user.Email == "" {
		responsesMap["email"] = "Email cannot be Empty"
	}

	if user.Password == "" {
		responsesMap["password"] = "Password cannot be Empty"
	}

	if user.FullName == "" {
		responsesMap["fullName"] = "Password cannot be Empty"
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

			responsesMap["status"] = "true"
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}
