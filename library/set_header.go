package library 

import (
	"net/http"
)

/*SetDefaultHTTPHeader ...
@desc set default HTTP Header to response
*/
func SetDefaultHTTPHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-auth-token")
}