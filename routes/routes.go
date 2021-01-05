package routes

import (
	"github.com/gorilla/mux"
	handler "fun-blogger-backend/handler"
	"net/http"
)

/*
	Register your routes here
*/
func Routing() http.Handler {
	router := mux.NewRouter()

	// Index Routes
	router.HandleFunc("/", handler. IndexHandler).Methods("GET", "POST", "PUT", "DELETE")

	// Register Routes
	router.HandleFunc("/register", handler. RegisterHandlerGet).Methods("GET")
	router.HandleFunc("/register", handler. RegisterHandlerPost).Methods("POST")

	return router
}