package middleware

import (
	library "fun-blogger-backend/library"
	"html"
	"net/http"
)

/*AuthMiddleware ...
@desc intercept user requests to check if token header is set
*/
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("x-auth-token")
		reqPath := html.EscapeString(r.URL.Path)

		// Which path required to using auth middleware
		var pathList = map[string]bool{
			"/login":    false,
			"/register": false,
			"/blog":     true,
		}

		if reqToken == "" && pathList[reqPath] == true {
			library.ResponseByCode(401, w, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
