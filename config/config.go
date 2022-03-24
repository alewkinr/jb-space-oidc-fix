package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config описывает структуру конфига
type Config struct {
	ServerHost   string `envconfig:"SERVER_HOST"`
	ServerPort   string `envconfig:"SERVER_PORT"`
	ProviderHost string `envconfig:"PROVIDER_HOST"`
}

// InitConfig возвращает конфиг
func InitConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	return &cfg, err
}

// MustInitConfig возвращает конфиг или паникует при ошибке
func MustInitConfig() *Config {
	cfg, err := InitConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
