package login

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/controllers"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	state := base64.StdEncoding.EncodeToString(b)
	return state, nil
}

func ServeLoginPage(w http.ResponseWriter, r *http.Request, auth *controllers.Authenticator) {
	state, err := generateRandomState()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "state",
		Value:    state,
		MaxAge:   300,
		Secure:   false,
		HttpOnly: true,
	})
	http.Redirect(w, r, auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
}
