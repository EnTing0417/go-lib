package model

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		URI string `yaml:"uri"`
	} `yaml:"database"`
	GoogleOAuth struct {
		ClientID     string   `yaml:"client_id"`
		ClientSecret string   `yaml:"client_secret"`
		RedirectUrl  string   `yaml:"redirect_url"`
		Scopes       []string `yaml:"scopes"`
		AuthURL      string   `yaml:"auth_url"`
		TokenURL     string   `yaml:"token_url"`
		LoginURL     string   `yaml:"login_url"`
	} `yaml:"googleOAuth"`
	Auth struct {
		PrivateKeyPemFile           string `yaml:"private_key_pem_file"`
		TkPublicKey            string `yaml:"tk_public_key"`
		RefreshTokenPrivateKey string `yaml:"refresh_tk_private_key_pem_file"`
		RefreshTokenPublicKey  string `yaml:"refresh_tk_public_key"`
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
