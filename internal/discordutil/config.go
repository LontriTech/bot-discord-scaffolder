package discordutil

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(filePath string) (*Config, error) {
	yamlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer yamlFile.Close()

	byteValue, err := io.ReadAll(yamlFile)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
