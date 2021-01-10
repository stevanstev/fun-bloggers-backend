package library 

import (
	"net/http"
)

/*SetDefaultHTTPHeader ...
@desc set default HTTP Header to response
*/
func SetDefaultHTTPHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set(
    	"Access-Control-Allow-Headers", 
    	"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-auth-token")
}