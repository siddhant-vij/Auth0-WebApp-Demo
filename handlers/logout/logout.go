package logout

import (
	"net/http"
	"net/url"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
	"github.com/siddhant-vij/Auth0-WebApp-Demo/utils"
)

func HandleLogout(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	logoutUrl, err := url.Parse("https://" + cfg.Domain + "/v2/logout")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	params := url.Values{}
	params.Add("returnTo", returnTo.String())
	params.Add("client_id", cfg.ClientID)
	logoutUrl.RawQuery = params.Encode()

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		MaxAge:   0,
		Secure:   false,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "state",
		Value:    "",
		MaxAge:   0,
		Secure:   false,
		HttpOnly: true,
	})

	cfg.UserProfile.Nickname = ""
	cfg.UserProfile.Picture = ""

	http.Redirect(w, r, logoutUrl.String(), http.StatusSeeOther)
}
