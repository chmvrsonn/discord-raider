package utils

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Message   string `json:"message"`
	ChannelID string `json:"channel_id"`
	TTS       bool   `json:"tts"`
	Delay     int    `json:"delay"`
}

func LoadConfig() (*Config, error) {
	if !FileExists("config.json") {
		defaultConfig := &Config{
			ChannelID: "",
			TTS:       false,
			Delay:     1000,
			Message:   "Default message",
		}

		if err := createDefaultConfigFile("config.json", defaultConfig); err != nil {
			return nil, err
		}
	}

	file, err := ReadRawFile("config.json")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func ValidateConfig(config *Config) error {
	if config.ChannelID == "" {
		return errors.New("channel ID is required")
	}

	if config.Delay < 0 {
		return errors.New("delay must be non-negative")
	}

	if config.Message == "" {
		return errors.New("message is required")
	}

	return nil
}

func createDefaultConfigFile(filename string, config *Config) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // For pretty-printing
	return encoder.Encode(config)
}
