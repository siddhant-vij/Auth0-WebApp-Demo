package home

import (
	"html/template"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func ServeHomePage(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	utils.GenerateHtml(tpl, "home", nil)
	http.ServeFile(w, r, "public/home.html")
}
