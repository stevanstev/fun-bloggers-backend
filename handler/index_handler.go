package handler

import (
	"encoding/json"
	"fmt"
	library "fun-blogger-backend/library"
	models "fun-blogger-backend/model"
	"net/http"
)

/*IndexHandler ...
@desc handling get request of /
@route /
@method GET
@access Public
*/
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(&w)

	response := models.BaseResponse{}
	response.GetDefault("Index Api Ready")

	httpResponse, err := json.Marshal(response)

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	fmt.Fprint(w, string(httpResponse))
}
