package main

import (
	"log"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/controllers"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/middlewares"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/router"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func main() {
	cfg := &config.Config{}
	auth := &controllers.Authenticator{}
	config.LoadEnv(cfg)
	cfg.SessionTokenMap = make(map[string]*oauth2.Token)
	cfg.PkceCodeVerifier = oauth2.GenerateVerifier()
	auth, err := controllers.NewAuthenticator(cfg)
	if err != nil {
		panic(err)
	}
	err = utils.CopyFiles("static", "public")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	corsMux := middlewares.CorsMiddleware(mux)
	router.RegisterRoutes(mux, cfg, auth)

	log.Fatal(http.ListenAndServe(":3000", corsMux))
}
