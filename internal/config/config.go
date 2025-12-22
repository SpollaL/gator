package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	data, err := io.ReadAll(configFile)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(data, &config)
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
	return filepath.Join(homeDir, configFileName), nil
}

func (config *Config) SetUser(userName string) error {
	config.CurrentUserName = userName

	configData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, configData, 0o644)
	if err != nil {
		return err
	}
	return nil
}
