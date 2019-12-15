package gonfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	typeJSON = "json"
)

func Load(path string, config interface{}) error {
	configFile, err := os.Open(path)
	defer configFile.Close()

	if err != nil {
		config = nil
		return LoadError
	}

	b, err := ioutil.ReadAll(configFile)
	if err != nil {
		config = nil
		return LoadError
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		config = nil
		return LoadError
	}
	return nil
}
