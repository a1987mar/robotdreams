package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Serv ServConfig
}

type ServConfig struct {
	Port string
	DNS  string
}

func LoadConfig() *Config {
	path := "../.env"
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Config{
		Serv: ServConfig{
			Port: os.Getenv("PORT"),
		},
	}
}
