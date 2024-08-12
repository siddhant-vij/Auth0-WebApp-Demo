package router

import (
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/handlers/home"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/handlers/user"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/middlewares"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

var cfg *config.Config = &config.Config{}

func init() {
	config.LoadEnv(cfg)
	err := utils.CopyFiles("static", "public")
	if err != nil {
		panic(err)
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.HandleFunc("/", home.ServeHomePage)
	mux.Handle("/user", middlewares.IsAuthenticated(http.HandlerFunc(user.ServeUserPage), cfg))
}
