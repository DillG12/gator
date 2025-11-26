package config

import (
	"encoding/json"
	"os"
)

func createFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := home + "/" + configFileName
	return filePath, nil
}

func Read() (Config, error) {
	filePath, err := createFilePath()
	if err != nil {
		return Config{}, err
	}
	//fmt.Println("Reading config from:", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	filePath, err := createFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
