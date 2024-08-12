package router

import (
	"html/template"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

var cfg *config.Config = &config.Config{}
var tpl *template.Template

func init() {
	config.LoadEnv(cfg)
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	utils.GenerateHtml(tpl, "home", nil)

	http.ServeFile(w, r, "public/home.html")
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", ServeHome)
}
