package router

import (
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
)

var cfg *config.Config = &config.Config{}

func init() {
	config.LoadEnv(cfg)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HelloWorld)
}
