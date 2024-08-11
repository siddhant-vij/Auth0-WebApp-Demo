package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(config *Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ClientID = os.Getenv("AUTH0_CLIENT_ID")
	config.Domain = os.Getenv("AUTH0_DOMAIN")
	config.ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	config.CallbackURL = os.Getenv("AUTH0_CALLBACK_URL")
}
