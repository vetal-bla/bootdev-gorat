package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err

	}

	return home + "/" + configFileName, nil
}

func Read() (Config, error) {

	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	c := Config{}

	if err := json.Unmarshal(byteValue, &c); err != nil {
		fmt.Println("Cant parse json")
		return Config{}, err
	}

	return c, nil
}

func write(cfg *Config) error {

	jsonString, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("Cant marshal json to byte")
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Cant open file")
		return err
	}

	err = os.WriteFile(filePath, jsonString, os.ModePerm)
	if err != nil {
		fmt.Println("Cant write file")
		return err
	}

	return nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(cfg)
}
