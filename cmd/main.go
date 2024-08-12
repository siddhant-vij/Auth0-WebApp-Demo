package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/middlewares"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/router"
)

func main() {
	mux := http.NewServeMux()
	corsMux := middlewares.CorsMiddleware(mux)
	router.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe("localhost:3000", corsMux))
}
