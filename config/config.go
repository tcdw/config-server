package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Token        string  `json:"token"`
	TemplatePath string  `json:"templatePath"`
	Port         float64 `json:"port"`
	Debug        bool    `json:"debug"`
}

func GetConfig(p string) (config *Config, err error) {
	configPath, err := filepath.Abs(p)
	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &config)

	return config, err
}
