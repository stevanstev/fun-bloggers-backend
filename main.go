package main

import (
	"fmt"
	routes "fun-blogger-backend/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func serve(mux http.Handler) {
	fmt.Println("Server Running on Port: 8080")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), mux))
	log.Fatal(http.ListenAndServe(":8009", handlers.CORS(originsOk, headersOk, methodsOk)(mux)))
}

func main() {
	mux := routes.Routing()
	serve(mux)
}
