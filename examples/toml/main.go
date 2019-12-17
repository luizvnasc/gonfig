package main

import (
	"fmt"

	"github.com/luizvnasc/gonfig"
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
	gonfig.Load("config.toml", &config)
	fmt.Printf("%v", config)
}
