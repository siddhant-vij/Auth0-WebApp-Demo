package user

import (
	"html/template"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func ServeUserPage(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/user.gohtml"))
	user := r.Context().Value(config.Config{}).(config.Config).UserProfile
	utils.GenerateHtml(tpl, "user", user)
	http.ServeFile(w, r, "public/user.html")
}
