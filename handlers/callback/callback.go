package callback

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/controllers"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func generateSessionId() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func ServeCallbackPage(w http.ResponseWriter, r *http.Request, auth *controllers.Authenticator, cfg *config.Config) {
	if r.URL.Query().Get("state") != cfg.SessionState {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid state parameter.")
		return
	}

	// Exchange an authorization code for a token.
	token, err := auth.Exchange(
		r.Context(),
		r.URL.Query().Get("code"),
		oauth2.VerifierOption(cfg.PkceCodeVerifier),
	)
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

	sessionId, err := generateSessionId()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	cfg.SessionTokenMap[sessionId] = token

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		MaxAge:   36000, // same as Access Token lifetime (Auth0)
		Secure:   false,
		HttpOnly: true,
	})
	http.Redirect(w, r, "/user", http.StatusTemporaryRedirect)
}
