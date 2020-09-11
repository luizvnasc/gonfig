package main

import (
	"fmt"

	"github.com/luizvnasc/gonfig/v2"
)

//Config struct to store the app configuration
type Config struct {
	Version     string `toml:"version"`
	Description string `toml:"desc"`
	Redis       struct {
		Host string `toml:"host"`
		Port uint   `toml:"port"`
	} `toml:"redis"`
}

func main() {
	var config Config
	gonfig.Load(&config, "config.toml")
	fmt.Printf("%v", config)
}
