package main

import (
	"fmt"

	"github.com/luizvnasc/gonfig"
)

//Config struct to store the app configuration
type Config struct {
	Version     string `env:"VERSION"`
	Description string `env:"DESCRIPTION"`
	Redis       struct {
		Host string `env:"REDIS_HOST"`
		Port uint   `env:"REDIS_PORT"`
	}
}

func main() {
	var config Config
	gonfig.Load(&config)
	fmt.Printf("%v", config)
}
