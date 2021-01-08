package routes

import (
	handler "fun-blogger-backend/handler"
	middleware "fun-blogger-backend/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

/*Routing ...
@desc register your routes here
*/
func Routing() http.Handler {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.AuthMiddleware)
	// Index Routes
	router.HandleFunc("/", handler.IndexHandler).Methods("GET", "POST", "PUT", "DELETE")

	// Register Routes
	router.HandleFunc("/register", handler.RegisterHandlerGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/register", handler.RegisterHandlerPost).Methods("POST", "OPTIONS")

	// Login Routes
	router.HandleFunc("/login", handler.LoginHandlerGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/login", handler.LoginHandlerPost).Methods("POST", "OPTIONS")

	// Blog Routes
	router.HandleFunc("/blog", handler.BlogHandlerGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/blog", handler.BlogHandlerPost).Methods("POST", "OPTIONS")

	// Relations Routes
	router.HandleFunc("/relations", handler.RelationsHandlerGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/relations/followed", handler.RelationsFollowedHandlerGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/relations/blocked", handler.RelationsBlockedHandlerGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/relations/followers", handler.RelationsByEmailHandlerPost).Methods("POST", "OPTIONS")
	router.HandleFunc("/relations/block", handler.RelationsBlockingUserHandlerPost).Methods("POST", "OPTIONS")
	router.HandleFunc("/relations/common", handler.FindCommonFollowersHandlerPost).Methods("POST", "OPTIONS")
	router.HandleFunc("/relations/follow", handler.RelationsFollowingUserHandlerPost).Methods("POST", "OPTIONS")

	return router
}
