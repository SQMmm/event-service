package config

import (
	"fmt"
	"github.com/subosito/gotenv"
	"os"
)

type MongoDB struct {
	Host string `env:"MONGO_DB_HOST"`
	Port string `env:"MONGO_DB_PORT"`
	DB   string `env:"MONGO_DB_NAME"`
}
type HTTP struct {
	Host string `env:"HTTP_HOST"`
	Port string `env:"HTTP_PORT"`
}

type Config struct {
	HTTP
	MongoDB
}

func Parse() (*Config, error) {
	err := gotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to get env: %s", err)
	}

	conf := &Config{
		HTTP: HTTP{
			Host: os.Getenv("HTTP_HOST"),
			Port: os.Getenv("HTTP_PORT"),
		},
		MongoDB: MongoDB{
			Host: os.Getenv("MONGO_DB_HOST"),
			Port: os.Getenv("MONGO_DB_PORT"),
			DB: os.Getenv("MONGO_DB_NAME"),
		},
	}

	return conf, nil
}
