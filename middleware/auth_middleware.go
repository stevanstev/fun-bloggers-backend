package middleware

import (
	driver "fun-blogger-backend/driver"
	library "fun-blogger-backend/library"
	"html"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*AuthMiddleware ...
@desc intercept user requests to check if token header is set
*/
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		library.SetDefaultHTTPHeader(w)

		reqToken := r.Header.Get("x-auth-token")
		reqPath := html.EscapeString(r.URL.Path)

		// Which path required to using auth middleware
		var pathList = map[string]bool{
			"/login":               false,
			"/register":            false,
			"/":                    false,
			"/blog":                true,
			"/relations":           true,
			"/relations/followed":  true,
			"/relations/blocked":   true,
			"/relations/followers": true,
			"/relations/block":     true,
			"/token/remove":     true,
			"/user/details":     true,
			"/user":     true,
		}

		// token is set and the pathList is match
		if reqToken == "" && pathList[reqPath] == true {
			library.ResponseByCode(401, w, "Unauthorized")
			return
		}

		if reqToken != "" {
			// Avoid sending random token
			userID := driver.GetUserIDByToken(reqToken)
			if userID == primitive.NilObjectID {
				library.ResponseByCode(401, w, "Unauthorized")
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
