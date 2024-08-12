package router

import (
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

var cfg *config.Config = &config.Config{}

func init() {
	config.LoadEnv(cfg)
}

func RandomResponseGen(w http.ResponseWriter, r *http.Request) {
	randomInt, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		panic(err)
	}

	if randomInt.Int64() == 0 {
		utils.RespondWithError(w, 400, "Goodbye World!")
	} else {
		utils.RespondWithJSON(w, http.StatusOK, "Hello World!")
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", RandomResponseGen)
}
