package user

import (
	"html/template"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func ServeUserPage(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	tpl := template.Must(template.ParseFiles("templates/user.gohtml"))
	utils.GenerateHtml(tpl, "user", cfg.UserProfile)
	http.ServeFile(w, r, "public/user.html")
}
