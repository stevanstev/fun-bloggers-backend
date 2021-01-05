package library

import (
	"net/http"
	"fmt"
)

/*
	API Response based on http status code
	@errorCode int, error code , 200, 500, ..
	@w http.ResponseWriter, the response writer
	@errorMessage string, error messages to send back to user
*/
func ResponseByCode(errorCode int, w http.ResponseWriter, errorMessage string) {
	switch errorCode {
		case 200:
			fmt.Fprint(w, errorMessage)
		case 500:
			http.Error(w, errorMessage, http.StatusInternalServerError)
		default:
			http.Error(w, errorMessage, http.StatusInternalServerError)
	}
}