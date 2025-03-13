package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Configs struct {
	TelegramToken    string `yaml:"TELEGRAM_BOT_TOKEN"`
	TelegramTokenEnv string `env:"TELEGRAM_BOT_TOKEN"`
}

func ExtractTelegramToken() (string, error) {
	var config Configs

	config.TelegramTokenEnv = os.Getenv("TELEGRAM_BOT_TOKEN")
	if config.TelegramTokenEnv != "" {
		return config.TelegramTokenEnv, nil
	}

	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return "", err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return "", err
	}

	return config.TelegramToken, nil
}
