package main

import (
	"log"
	"net/http"
	"user-app/routes"
)

func main() {
	r := routes.SetupRoutes()
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
