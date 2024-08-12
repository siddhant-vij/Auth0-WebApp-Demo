package middlewares

import (
	"context"
	"net/http"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
)

func IsAuthenticated(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.UserProfile == (config.Config{}).UserProfile {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), config.Config{}, *cfg))
		next.ServeHTTP(w, r)
	})
}
