package gonfig

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
)

type unmarshalerFunc func(data []byte, v interface{}) error

// Load a struct with the configuration from a config file.
func Load(path string, config interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		config = nil
		return LoadError
	}
	defer configFile.Close()

	b, err := ioutil.ReadAll(configFile)
	if err != nil {
		config = nil
		return LoadError
	}
	unmarshaler, err := getUnmarchaler(path)
	if err != nil {
		return err
	}
	err = unmarshaler(b, &config)
	if err != nil {
		config = nil
		return LoadError
	}
	return nil
}

func getUnmarchaler(path string) (unmarshalerFunc, error) {
	ext := filepath.Ext(path)

	switch ext {
	case ".json":
		return json.Unmarshal, nil
	case ".xml":
		return xml.Unmarshal, nil
	default:
		return nil, UnsupportedFileError
	}
}
