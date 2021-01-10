package handler

import (
	"net/http"
	library "fun-blogger-backend/library"
	driver "fun-blogger-backend/driver"
	"go.mongodb.org/mongo-driver/bson"
	"encoding/json"
)

/*UserGetDetailsByToken ...
@desc handling get request of /user/details
@route /user/details
@method GET
@access Private
*/
func UserGetDetailsByToken(w http.ResponseWriter, r * http.Request) {
	library.SetDefaultHTTPHeader(&w)

	reqToken := r.Header.Get("x-auth-token")
	userID := driver.GetUserIDByToken(reqToken)

	var responsesMap = make(map[string]interface{}, 0)

	query := bson.M{
		"_id": userID,
	}

	userDetails, _ := driver.FindUsers(query)

	responsesMap["details"] = userDetails
	responsesMap["status"] = "true"

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}