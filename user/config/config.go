package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	Server struct {
		Host string `env:"SERVER_HOST"`
		Port string `env:"SERVER_PORT"`
	}
	Pg struct {
		User     string `env:"PG_USER"`
		Host     string `env:"PG_HOST"`
		Database string `env:"PG_DATABASE"`
		Password string `env:"PG_PASSWORD"`
		Port     string `env:"PG_PORT"`
	}
}

func New() *Config {
	c := &Config{}
	file, err := os.Open("./config.env")
	err = dotenv.NewDecoder(file).Decode(c)
	if err != nil {
		panic(err)
	}
	return c
}
