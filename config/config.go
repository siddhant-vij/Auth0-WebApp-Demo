package config

type Config struct {
	ClientID     string
	Domain       string
	ClientSecret string
	CallbackURL  string

	UserProfile struct {
		Nickname string
		Picture  string
	}
}
