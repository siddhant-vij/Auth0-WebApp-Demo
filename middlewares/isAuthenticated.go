package middlewares

import (
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
)

func IsAuthenticated(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.UserProfile == (config.Config{}).UserProfile {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
