package gonfig

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/luizvnasc/goenv"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type unmarshalerFunc func(data []byte, v interface{}) error

// Load a struct with the configuration from a config file.
func Load(config interface{}, args ...string) error {
	if len(args) == 0 {
		err := goenv.Unmarshal(config)
		return err
	}
	path := args[0]
	unmarshaler, err := getUnmarshaler(path)
	if err != nil {
		return err
	}
	content, err := getFileContent(path)
	if err != nil {
		return LoadError
	}
	err = unmarshaler(content, config)
	if err != nil {
		return LoadError
	}
	return nil
}

func getUnmarshaler(path string) (unmarshalerFunc, error) {
	ext := filepath.Ext(path)

	switch {
	case ext == ".json":
		return json.Unmarshal, nil
	case ext == ".xml":
		return xml.Unmarshal, nil
	case ext == ".yaml" || ext == ".yml":
		return yaml.Unmarshal, nil
	case ext == ".toml":
		return toml.Unmarshal, nil
	default:
		return nil, UnsupportedFileError
	}
}

func getFileContent(path string) (bytes []byte, err error) {
	configFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer configFile.Close()
	bytes, err = ioutil.ReadAll(configFile)
	return
}
