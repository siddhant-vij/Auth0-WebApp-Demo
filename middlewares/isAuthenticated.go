package middlewares

import (
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/controllers"
)

func IsAuthenticated(next http.Handler, auth *controllers.Authenticator, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.UserProfile == (config.Config{}).UserProfile {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		sessionId, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		token, ok := cfg.SessionTokenMap[sessionId.Value]
		if !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		_, err = auth.VerifyIDToken(r.Context(), token)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
