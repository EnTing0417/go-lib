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
	} `yaml:"googleOAuth"`
	Auth struct {
		SecretKey string `yaml:"secret_key"`
		RefreshTokenSecretKey string `yaml:"refresh_token_secret_key"`
	} `yaml:"auth"`
}

func ReadConfig() (config *Config) {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
		return nil
	}

	return config
}