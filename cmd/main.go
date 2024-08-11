package main

import (
	"fmt"

	"github.com/siddhant-vij/Auth0-WebApp-Demo/config"
)

var cfg *config.Config = &config.Config{}

func init() {
	config.LoadEnv(cfg)
}

func main() {
	fmt.Println("Checking if config properly loaded")
	fmt.Println(cfg)
}
