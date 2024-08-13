package router

import (
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/controllers"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/handlers/callback"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/handlers/home"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/handlers/login"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/handlers/logout"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/handlers/user"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/middlewares"
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

func RegisterRoutes(mux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.HandleFunc("/", home.ServeHomePage)

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		login.ServeLoginPage(w, r, auth)
	})

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		callback.ServeCallbackPage(w, r, auth, cfg)
	})

	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		logout.HandleLogout(w, r, cfg)
	})

	mux.Handle("/user", middlewares.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user.ServeUserPage(w, r, cfg)
	}), cfg))
}
