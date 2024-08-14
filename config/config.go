package config

import (
	"golang.org/x/oauth2"
)

type Config struct {
	ClientID     string
	Domain       string
	ClientSecret string
	CallbackURL  string

	SessionState    string
	SessionTokenMap map[string]*oauth2.Token

	UserProfile struct {
		Nickname string
		Picture  string
	}
}
