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
	validJsonFile       = "./test/config.json"
	invalidJsonFile     = "./test/invalid.json"
	invalidJsonBodyFile = "./test/invalid_config.json"
	validXMLFile        = "./test/config.xml"
	invalidXMLFile      = "./test/invalid.xml"
	invalidXMLBodyFile  = "./test/invalid_config.xml"
	unsupportedFile     = "./test/config.xyz"
)

type SomeConfiguration struct {
	Version     string `json:"version" xml:"version"`
	ProjectName string `json:"project_name" xml:"project-name"`
}

var configValid SomeConfiguration

func init() {
	file, err := os.Open(validJsonFile)
	if err != nil {
		panic("Error Loading teste sample")
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Error Loading teste sample")
	}
	json.Unmarshal(b, &configValid)
}
func TestGonfig(t *testing.T) {
	config := SomeConfiguration{}
	t.Run("JSON tests", func(t *testing.T) {
		t.Run("Load a configuration from a valid json file", func(t *testing.T) {
			err := gonfig.Load(validJsonFile, &config)
			if err != nil {
				t.Errorf("Error loading the configuration: %v", err)
			}
			if !reflect.DeepEqual(config, configValid) {
				t.Errorf("Error loading the configuration: expected %v, got %v", configValid, config)
			}
		})

		t.Run("Load a configuration from an invalid json file", func(t *testing.T) {
			err := gonfig.Load(invalidJsonFile, &config)
			if err == nil {
				t.Errorf("It was expected to get an error. Got nil")
			}
			if err != gonfig.LoadError {
				t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
			}
		})

		t.Run("Load a configuration from an invalid json body", func(t *testing.T) {
			err := gonfig.Load(invalidJsonBodyFile, &config)
			if err == nil {
				t.Errorf("It was expected to get an error. Got nil")
			}
			if err != gonfig.LoadError {
				t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
			}
		})
	})
	t.Run("XML tests", func(t *testing.T) {
		t.Run("Load a configuration from a valid xml file", func(t *testing.T) {
			err := gonfig.Load(validXMLFile, &config)
			if err != nil {
				t.Errorf("Error loading the configuration: %v", err)
			}
			if !reflect.DeepEqual(config, configValid) {
				t.Errorf("Error loading the configuration: expected %v, got %v", configValid, config)
			}
		})
		t.Run("Load a configuration from an invalid json file", func(t *testing.T) {
			err := gonfig.Load(invalidXMLFile, &config)
			if err == nil {
				t.Errorf("It was expected to get an error. Got nil")
			}
			if err != gonfig.LoadError {
				t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
			}
		})

		t.Run("Load a configuration from an invalid json body", func(t *testing.T) {
			err := gonfig.Load(invalidXMLBodyFile, &config)
			if err == nil {
				t.Errorf("It was expected to get an error. Got nil")
			}
			if err != gonfig.LoadError {
				t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
			}
		})
	})
	t.Run("Unsupported file", func(t *testing.T) {
		err := gonfig.Load(unsupportedFile, &config)
		if err == nil {
			t.Errorf("It was expected to get an error. Got nil")
		}
		if err != gonfig.UnsupportedFileError {
			t.Errorf("Expected the error %v, got %v", gonfig.UnsupportedFileError, err)
		}
	})

}
