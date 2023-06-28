package model

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Database struct {
		URI     string `yaml:"uri"`
	} `yaml:"database"`
	GoogleOAuth struct {
		ClientID     string `yaml:"client_id"`
		ClientSecret     string `yaml:"client_secret"`
		RedirectUrl string `yaml:"redirect_url"`
		Scopes	[]string `yaml:"scopes"`
		AuthURL	string `yaml:"auth_url"`
		TokenURL	string `yaml:"token_url"`
		LoginURL	string `yaml:"login_url"`
		TokenExpiresIn int `yaml:"token_expires_in"`
	} `yaml:"googleOAuth"`
}

func ReadConfig() (config *Config) {
	// Read the config.yaml file
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil
	}

	// Parse the YAML data into the Config struct
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
		return nil
	}

	return config
}