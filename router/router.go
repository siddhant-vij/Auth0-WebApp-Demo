package router

import (
	"crypto/rand"
	"html/template"
	"math/big"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
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

func ServeRandomPage(w http.ResponseWriter, r *http.Request) {
	randomInt, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		panic(err)
	}

	if randomInt.Int64() == 0 {
		tpl := template.Must(template.ParseFiles("templates/home.gohtml"))
		utils.GenerateHtml(tpl, "home", nil)
		http.ServeFile(w, r, "public/home.html")
	} else {
		tpl := template.Must(template.ParseFiles("templates/user.gohtml"))
		user := struct {
			Picture, Nickname string
		}{
			Picture:  "https://i.imgur.com/Y5Q2C0N.png",
			Nickname: "Siddhant",
		}
		utils.GenerateHtml(tpl, "user", user)
		http.ServeFile(w, r, "public/user.html")
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.HandleFunc("/", ServeRandomPage)
}
