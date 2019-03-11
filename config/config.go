package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	OAuth2 map[string]struct {
		ClientID     string   `yaml:"client_id"`
		ClientSecret string   `yaml:"client_secret"`
		RedirectURL  string   `yaml:"redirect_url"`
		Scopes       []string `yaml:"scopes"`
	} `yaml:"oauth2"`
}

const configYaml = "config.yaml"

var config *Config

// Load loads configuration
func Load() {
	b, err := ioutil.ReadFile(configYaml)
	if err != nil {
		log.Fatalf("Could not read the %s: %v", configYaml, err)
	}
	if err = yaml.Unmarshal(b, &config); err != nil {
		log.Fatalf("Could not parse the %s: %v", configYaml, err)
	}
}

// Get returns global configuration
func Get() *Config {
	if config == nil {
		Load()
	}
	return config
}
