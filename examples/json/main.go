package main

import (
	"fmt"

	"github.com/luizvnasc/gonfig"
)

//Config struct to store the app configuration
type Config struct {
	Version     string `json:"version"`
	Description string `json:"desc"`
	Redis       struct {
		Host string `json:"host"`
		Port uint   `json:"port"`
	} `json:"redis"`
}

func main() {
	var config Config
	gonfig.Load(&config, "config.json")
	fmt.Printf("%v", config)
}
