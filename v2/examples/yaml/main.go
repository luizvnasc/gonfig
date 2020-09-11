package main

import (
	"fmt"

	"github.com/luizvnasc/gonfig/v2"
)

//Config struct to store the app configuration
type Config struct {
	Version     string `yaml:"version"`
	Description string `yaml:"desc"`
	Redis       struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"redis"`
}

func main() {
	var config Config
	gonfig.Load(&config, "config.yaml")
	fmt.Printf("%v", config)
}
