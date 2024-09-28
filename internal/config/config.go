package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"
type Config struct{
	DBURL string						`json:"db_url"`
	CurrentUserName string	`json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}
	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	err := write(*cfg)
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