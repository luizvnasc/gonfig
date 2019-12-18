package main

import (
	"fmt"

	"github.com/luizvnasc/gonfig"
)

//Config struct to store the app configuration
type Config struct {
	Version     string `xml:"version"`
	Description string `xml:"description"`
	Redis       struct {
		Host string `xml:"host,attr"`
		Port uint   `xml:"port,attr"`
	} `xml:"redis"`
}

func main() {
	var config Config
	gonfig.Load(&config, "config.xml")
	fmt.Printf("%v", config)
}
