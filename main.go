package main 

import (
	"fmt"
	"net/http"
	"log"
	routes "fun-blogger-backend/routes"
)

func serve(mux http.Handler) {
	fmt.Println("Server Running on Port: 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}

func main() {
	mux := routes.Routing();
	serve(mux)
}