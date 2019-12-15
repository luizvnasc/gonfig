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

	ext := filepath.Ext(path)

	var unmarshaler unmarshalerFunc
	switch ext {
	case ".json":
		unmarshaler = json.Unmarshal
	case ".xml":
		unmarshaler = xml.Unmarshal
	default:
		return UnsupportedFileError
	}
	err = unmarshaler(b, &config)
	if err != nil {
		config = nil
		return LoadError
	}
	return nil
}
