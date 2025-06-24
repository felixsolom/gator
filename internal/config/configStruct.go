package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) *Config {
	c.CurrentUserName = user
	return c
}

func Read(fPath string, config *Config) (*Config, error) {
	fullPath, err := FilePath(fPath)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("could't read from json file: %v", err)
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal json file: %v", err)
	}
	return config, nil
}

func Write(fPath string, config *Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("couldn't convert to json data %v", err)
	}
	fullPath, err := FilePath(fPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(fullPath, data, 0644)
	if err != nil {
		return fmt.Errorf("couldn't write to json file %v", err)
	}
	return nil
}
