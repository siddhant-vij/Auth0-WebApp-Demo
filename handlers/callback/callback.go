package callback

import (
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/controllers"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func ServeCallbackPage(w http.ResponseWriter, r *http.Request, auth *controllers.Authenticator, cfg *config.Config) {
	stateCookie, err := r.Cookie("state")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Couldn't find state cookie.")
		return
	}
	if r.URL.Query().Get("state") != stateCookie.Value {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid state parameter.")
		return
	}

	// Exchange an authorization code for a token.
	token, err := auth.Exchange(r.Context(), r.URL.Query().Get("code"))
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
		return
	}

	idToken, err := auth.VerifyIDToken(r.Context(), token)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to verify ID Token.")
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to parse ID Token claims.")
		return
	}

	cfg.UserProfile.Nickname = profile["nickname"].(string)
	cfg.UserProfile.Picture = profile["picture"].(string)

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		MaxAge:   36000,
		Secure:   false,
		HttpOnly: true,
	})
	http.Redirect(w, r, "/user", http.StatusTemporaryRedirect)
}
