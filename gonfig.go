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
	content, err := getFileContent(path)
	if err != nil {
		return LoadError
	}
	unmarshaler, err := getUnmarshaler(path)
	if err != nil {
		return err
	}
	err = unmarshaler(content, &config)
	if err != nil {
		config = nil
		return LoadError
	}
	return nil
}

func getUnmarshaler(path string) (unmarshalerFunc, error) {
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

func getFileContent(path string) (bytes []byte, err error) {
	configFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer configFile.Close()
	bytes, err = ioutil.ReadAll(configFile)
	return
}
