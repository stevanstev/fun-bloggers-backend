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
	router.HandleFunc("/register", handler.RegisterHandlerGet).Methods("GET")
	router.HandleFunc("/register", handler.RegisterHandlerPost).Methods("POST")

	// Login Routes
	router.HandleFunc("/login", handler.LoginHandlerPost).Methods("POST")
	router.HandleFunc("/login", handler.LoginHandlerPost).Methods("GET")

	// Blog Routes
	router.HandleFunc("/blog", handler.BlogHandlerGet).Methods("GET")
	router.HandleFunc("/blog", handler.BlogHandlerPost).Methods("POST")

	// Relations Routes
	router.HandleFunc("/relations", handler.RelationsHandlerGet).Methods("GET")
	router.HandleFunc("/relations/followed", handler.RelationsFollowedHandlerGet).Methods("GET")
	router.HandleFunc("/relations/blocked", handler.RelationsBlockedHandlerGet).Methods("GET")

	return router
}
