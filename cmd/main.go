package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/middlewares"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/router"
)

func main() {
	log.Println("Checking if HelloWorld server is running...")
	mux := http.NewServeMux()
	corsMux := middlewares.CorsMiddleware(mux)
	router.RegisterRoutes(mux)

	log.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", corsMux)
	if err != nil {
		log.Fatal(err)
	}
}
