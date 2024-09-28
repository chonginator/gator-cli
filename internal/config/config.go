package config

import (
	"encoding/json"
	"os"
)

type Config struct{
	DBURL string						`json:"db_url"`
	CurrentUserName string	`json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return (
		homeDir +
		"/Documents/workspace/github.com/chonginator/gator-cli/" +
		configFileName), nil
}

func (cfg Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	err := write(cfg)
	if err != nil { 
		return err
	}
	return nil
}

func write(cfg Config) error {
	configFileName, err := getConfigFilePath()
	if err != nil {
		return err
	}

	configData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFileName, configData, 0644)
	if err != nil {
		return err
	}
	return nil
}