package login

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/controllers"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func ServeLoginPage(w http.ResponseWriter, r *http.Request, auth *controllers.Authenticator, cfg *config.Config) {
	state, err := generateRandomState()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	cfg.SessionState = state
	http.Redirect(w, r,
		auth.AuthCodeURL(
			state,
			oauth2.S256ChallengeOption(cfg.PkceCodeVerifier),
		),
		http.StatusTemporaryRedirect)
}
