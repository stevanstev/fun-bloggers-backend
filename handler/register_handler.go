package handler

import (
	"net/http"
	models "fun-blogger-backend/model"
	library "fun-blogger-backend/library"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	driver "fun-blogger-backend/driver"
)

func RegisterHandlerGet(w http.ResponseWriter, r * http.Request) {
	response := models.BaseResponse{}
	response.GetDefault("Register Api Ready")

	w.Header().Add("Content-Type", "application/json")
	httpResponse, err := json.Marshal(response)

	if (err != nil) {
		library.ResponseByCode(500, w, err.Error())
		return 
	}

	fmt.Fprint(w, string(httpResponse))
}

func RegisterHandlerPost(w http.ResponseWriter, r *http.Request) {
	var model models.User
	err := model.FromJson(r)

	var responsesMap = map[string]string{}

	if (err != nil) {
		library.ResponseByCode(500, w, err.Error())
		return 
	}

	var email string = model.Email 
	var password string = model.Password 
	var fullName string = model.FullName

	if (email == "") {
		responsesMap["email"] = "Email cannot be Empty" 
	}

	if (password == "") {
		responsesMap["password"] = "Password cannot be Empty" 
	}

	if (fullName == "") {
		responsesMap["fullName"] = "Password cannot be Empty" 
	}

	if (len(responsesMap) != 0) {
		responsesMap["status"] = "false"
	} else {
		// Set User ID
		model.Id = primitive.NewObjectID()
		model.CreatedAt = library.GetCurrentDate()
		model.UpdatedAt = library.GetCurrentDate()
		model.Password, _ = library.Encrypt([]byte(password))

		// Insert model to users table
		err = driver.Insert("users", model)

		if (err != nil) {
			library.ResponseByCode(500, w, "There's something wrong")
			return
		}

		responsesMap["status"] = "true"
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	
	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}
