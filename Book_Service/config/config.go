package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Admin    string `yaml:"admin"`
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
	} `yaml:"database"`

	Grpc struct {
		Port string `yaml:"port"`
	} `yaml:"grpc"`
}

func LoadConfig() (*Config, error) {
	// Добираемся до файла env
	data, err := os.ReadFile("config.env")
	if err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(data, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg

}
