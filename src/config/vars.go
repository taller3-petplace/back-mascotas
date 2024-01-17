package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const envFile = ".env"

type AppConfig struct {
	Port  int
	DbURL string
}

func LoadConfig() (AppConfig, error) {

	var config AppConfig
	if err := godotenv.Load(); err != nil {
		log.Print("error cargando el archivo: ", err)
	}

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if portStr == "" || err != nil {
		config.Port = 9000
	} else {
		config.Port = port
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return config, errors.New("missing url")
	}

	config.DbURL = dbUrl

	return config, nil
}
