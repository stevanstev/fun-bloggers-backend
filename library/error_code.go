package library

import (
	"fmt"
	"net/http"
)

/*ResponseByCode ...
@desc API Response based on http status code
@param errorCode int, error code , 200, 500, ..
@param w http.ResponseWriter, the response writer
@param errorMessage string, error messages to send back to user
*/
func ResponseByCode(errorCode int, w http.ResponseWriter, errorMessage string) {
	switch errorCode {
	case 200:
		// w will be closing to avoid memory leaked
		fmt.Fprint(w, errorMessage)
	case 401:
		fmt.Fprint(w, errorMessage)
	case 500:
		http.Error(w, errorMessage, http.StatusInternalServerError)
	default:
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}
}
