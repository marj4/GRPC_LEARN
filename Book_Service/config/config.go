package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Type     string `env:"DB_TYPE"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Name     string `env:"DB_NAME"`
	Password string `env:"DB_PASSWORD"`
	SslMode  string `env:"DB_SSL"`
}

func LoadConfig() (*Config, error) {
	// Путь до .env файла
	path := "C:/Users/MarJ/Desktop/GRPC_V2/Book_Service/config.env"

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg, nil

}
