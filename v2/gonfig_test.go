package gonfig_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/luizvnasc/gonfig/v2"
)

const (
	validJsonFile       = "./test/config.json"
	invalidJsonFile     = "./test/invalid.json"
	invalidJsonBodyFile = "./test/invalid_config.json"
	validXMLFile        = "./test/config.xml"
	invalidXMLBodyFile  = "./test/invalid_config.xml"
	unsupportedFile     = "./test/config.xyz"
	validYamlFile       = "./test/config.yaml"
	invalidYamlBodyFile = "./test/invalid_config.yaml"
	validTomlFile       = "./test/config.toml"
	invalidTomlBodyFile = "./test/invalid_config.toml"
)

type SomeConfiguration struct {
	Version     string `json:"version" xml:"version" yaml:"version" toml:"version" env:"VERSION"`
	ProjectName string `json:"project_name" xml:"project-name" yaml:"project_name" toml:"project_name" env:"PROJECT_NAME"`
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
	t.Run("JSON tests", func(t *testing.T) {
		t.Run("Load a configuration from a valid json file", func(t *testing.T) {
			config := SomeConfiguration{}
			err := gonfig.Load(&config, validJsonFile)
			if err != nil {
				t.Errorf("Error loading the configuration: %v", err)
			}
			if !reflect.DeepEqual(config, configValid) {
				t.Errorf("Error loading the configuration: expected %v, got %v", configValid, config)
			}
		})

		t.Run("Load a configuration from an invalid json file", func(t *testing.T) {
			config := SomeConfiguration{}
			err := gonfig.Load(&config, invalidJsonFile)
			if err == nil {
				t.Errorf("It was expected to get an error. Got nil")
			}
			if err != gonfig.LoadError {
				t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
			}
		})

		t.Run("Load a configuration from an invalid json body", func(t *testing.T) {
			config := SomeConfiguration{}
			err := gonfig.Load(&config, invalidJsonBodyFile)
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
			config := SomeConfiguration{}
			err := gonfig.Load(&config, validXMLFile)
			if err != nil {
				t.Errorf("Error loading the configuration: %v", err)
			}
			if !reflect.DeepEqual(config, configValid) {
				t.Errorf("Error loading the configuration: expected %v, got %v", configValid, config)
			}
		})
		t.Run("Load a configuration from an invalid xml body", func(t *testing.T) {
			config := SomeConfiguration{}
			err := gonfig.Load(&config, invalidXMLBodyFile)
			if err == nil {
				t.Errorf("It was expected to get an error. Got nil")
			}
			if err != gonfig.LoadError {
				t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
			}
		})
	})
	t.Run("YAML tests", func(t *testing.T) {
		t.Run("Load a configuration from a valid yaml file", func(t *testing.T) {
			config := SomeConfiguration{}
			err := gonfig.Load(&config, validYamlFile)
			if err != nil {
				t.Errorf("Error loading the configuration: %v", err)
			}
			if !reflect.DeepEqual(config, configValid) {
				t.Errorf("Error loading the configuration: expected %v, got %v", configValid, config)
			}
		})

		t.Run("Load a configuration from an invalid yaml body", func(t *testing.T) {
			config := SomeConfiguration{}
			err := gonfig.Load(&config, invalidYamlBodyFile)
			if err == nil {
				t.Errorf("It was expected to get an error. Got nil")
			}
			if err != gonfig.LoadError {
				t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
			}
		})
	})

	t.Run("TOML tests", func(t *testing.T) {
		t.Run("Load a configuration from a valid toml file", func(t *testing.T) {
			config := SomeConfiguration{}
			err := gonfig.Load(&config, validTomlFile)
			if err != nil {
				t.Errorf("Error loading the configuration: %v", err)
			}
			if !reflect.DeepEqual(config, configValid) {
				t.Errorf("Error loading the configuration: expected %v, got %v", configValid, config)
			}
		})

		t.Run("Load a configuration from an invalid toml body", func(t *testing.T) {
			config := SomeConfiguration{}
			err := gonfig.Load(&config, invalidTomlBodyFile)
			if err == nil {
				t.Errorf("It was expected to get an error. Got nil")
			}
			if err != gonfig.LoadError {
				t.Errorf("Expected the error %v, got %v", gonfig.LoadError, err)
			}
		})
	})

	t.Run("ENV tests", func(t *testing.T) {
		t.Run("Load a configuration from ENV", func(t *testing.T) {
			config := SomeConfiguration{}
			os.Setenv("VERSION", configValid.Version)
			os.Setenv("PROJECT_NAME", configValid.ProjectName)
			err := gonfig.Load(&config)
			if err != nil {
				t.Errorf("Error loading the configuration: %v", err)
			}
			if !reflect.DeepEqual(config, configValid) {
				t.Errorf("Error loading the configuration: expected %v, got %v", configValid, config)
			}
		})

	})

	t.Run("Unsupported file", func(t *testing.T) {
		config := SomeConfiguration{}
		err := gonfig.Load(&config, unsupportedFile)
		if err == nil {
			t.Errorf("It was expected to get an error. Got nil")
		}
		if err != gonfig.UnsupportedFileError {
			t.Errorf("Expected the error %v, got %v", gonfig.UnsupportedFileError, err)
		}
	})

}
