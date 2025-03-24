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
	port string
}

func LoadConfig() *Config {
	path := "../.env"
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Config{
		Serv: ServConfig{
			port: os.Getenv("PORT"),
		},
	}
}
