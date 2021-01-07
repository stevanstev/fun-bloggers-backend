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

/*LoginHandlerGet ...
@desc handling get request of /login
@route /login
@method GET
@access Public
*/
func LoginHandlerGet(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)

	response := models.BaseResponse{}
	response.GetDefault("Login Api Ready")

	w.Header().Add("Content-Type", "application/json")
	httpResponse, err := json.Marshal(response)

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	fmt.Fprint(w, string(httpResponse))
}

/*LoginHandlerPost ...
@desc handling post request of /login
@route /login
@method POST
@access Public
*/
func LoginHandlerPost(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if r.Body == nil {
		library.ResponseByCode(500, w, "No request body specified")
		return
	}

	err := user.FromJSON(r)

	var responsesMap = map[string]string{}

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if user.Email == "" {
		responsesMap["email"] = "Email cannot be empty"
	}

	if user.Password == "" {
		responsesMap["password"] = "Password cannot be empty"
	}

	if len(responsesMap) != 0 {
		responsesMap["status"] = "false"
	} else {
		query := bson.M{"email": user.Email}

		result, err := driver.FindUsers(query)

		if err != nil {
			library.ResponseByCode(500, w, err.Error())
			return
		}

		if len(result) == 0 {
			// If data is not exists
			responsesMap["email"] = "Email or Password is Wrong"
			responsesMap["status"] = "false"
		} else {
			// If Password doen't match
			err = library.CompareHashedPassword(result[0].Password, user.Password)

			if err != nil {
				responsesMap["password"] = "Email or Password is Wrong"
				responsesMap["status"] = "false"
			} else {
				//Send Token
				var token models.Token
				token.ID = primitive.NewObjectID()
				token.UserID = result[0].ID
				token.CreatedAt = library.GetCurrentDate()
				token.UpdatedAt = library.GetCurrentDate()

				jwtToken := library.Sign(result, token.UserID)

				token.Token = jwtToken

				err = driver.Insert("tokens", token)

				if err != nil {
					library.ResponseByCode(500, w, err.Error())
					return
				}

				responsesMap["token"] = jwtToken
				responsesMap["status"] = "true"
			}
		}
	}

	library.SetDefaultHTTPHeader(w)
	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}
