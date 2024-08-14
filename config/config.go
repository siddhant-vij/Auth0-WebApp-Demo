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

	PkceCodeVerifier string

	UserProfile struct {
		Name string
		Picture  string
	}
}
