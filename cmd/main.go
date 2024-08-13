package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/controllers"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/middlewares"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/router"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

var (
	cfg  *config.Config             = &config.Config{}
	auth *controllers.Authenticator = &controllers.Authenticator{}
)

func init() {
	config.LoadEnv(cfg)
	auth0, err := controllers.NewAuthenticator(cfg)
	if err != nil {
		panic(err)
	}
	auth = auth0
	err = utils.CopyFiles("static", "public")
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	corsMux := middlewares.CorsMiddleware(mux)
	router.RegisterRoutes(mux, cfg, auth)

	log.Fatal(http.ListenAndServe("localhost:3000", corsMux))
}
