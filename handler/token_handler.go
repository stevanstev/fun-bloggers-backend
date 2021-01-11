package handler

import (
	"encoding/json"
	driver "fun-blogger-backend/driver"
	library "fun-blogger-backend/library"
	"net/http"
)

/*TokenHandlerDeletePost ...
@desc handling post request of /token/remove
@route /token/remove
@method POST
@access Private
*/
func TokenHandlerDeletePost(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(&w)

	type tokenType struct {
		Token string `json:"token"`
	}

	var tokenString tokenType
	var responsesMap = make(map[string]interface{}, 0)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tokenString)

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if tokenString.Token == "" {
		responsesMap["tokenError"] = "token is not set"
	}

	if len(responsesMap) != 0 {
		responsesMap["status"] = "false"
	} else {
		error := driver.DeleteToken(tokenString.Token)

		if error != nil {
			responsesMap["status"] = "false"
		}

		responsesMap["status"] = "true"
	}

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}
