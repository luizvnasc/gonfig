package gonfig_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/luizvnasc/gonfig"
)

const (
	validJsonFile       = "config.json"
	invalidJsonFile     = "invalid.json"
	invalidJsonBodyFile = "invalid_config.json"
)

type SomeConfiguration struct {
	Version     string `json:"version" xml:"version"`
	ProjectName string `json:"project_name" xml:"project_name"`
}

var configJSON SomeConfiguration

func init() {
	file, err := os.Open(validJsonFile)
	if err != nil {
		panic("Error Loading teste sample")
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Error Loading teste sample")
	}
	json.Unmarshal(b, &configJSON)
}
func TestLoadJSON(t *testing.T) {
	config := SomeConfiguration{}

	t.Run("Load a configuration from a valid json file", func(t *testing.T) {
		err := gonfig.Load(validJsonFile, &config)
		if err != nil {
			t.Errorf("Error loading the configuration: %v", err)
		}
		if !reflect.DeepEqual(config, configJSON) {
			t.Errorf("Error loading the configuration: expected %v, got %v", configJSON, config)
		}
	})

	t.Run("Load a configuration from a invalid json file", func(t *testing.T) {
		err := gonfig.Load(invalidJsonFile, &config)
		if err == nil {
			t.Errorf("It was expected to get an error. Got nil")
		}
		if err != gonfig.LoadError {
			t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
		}
	})

	t.Run("Load a configuration from a invalid json body", func(t *testing.T) {
		err := gonfig.Load(invalidJsonFile, &config)
		if err == nil {
			t.Errorf("It was expected to get an error. Got nil")
		}
		if err != gonfig.LoadError {
			t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
		}
	})

}
