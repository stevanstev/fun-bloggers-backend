package main

import (
	"fmt"
	routes "fun-blogger-backend/routes"
	"log"
	"net/http"
	"os"
	// "github.com/gorilla/handlers"
)

func serve(mux http.Handler) {
	fmt.Println("Server Running")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), mux))
	// log.Fatal(http.ListenAndServe(":8009", mux))
}

func main() {
	mux := routes.Routing()
	serve(mux)
}
