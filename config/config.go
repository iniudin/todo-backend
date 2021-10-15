package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT"`
	}

	Database struct {
		Name     string `env:"DB_NAME"`
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		Username string `env:"DB_USER"`
		Password string `env:"DB_PASS"`
	}
}

func NewCOnfig() (*Config, error) {
	config := Config{}
	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}

	err = dotenv.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
