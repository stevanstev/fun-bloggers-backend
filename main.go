package main 

import (
	"fmt"
	"net/http"
	"log"
	"os"
	routes "fun-blogger-backend/routes"
)

func serve(mux http.Handler) {
	fmt.Println("Server Running on Port: 8080")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), mux))
}

func main() {
	mux := routes.Routing();
	serve(mux)
}