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
	Port         int
	DbURL        string
	TreatmentURL string
	LogLevel     string
}

func LoadConfig() (AppConfig, error) {

	var config AppConfig
	if err := godotenv.Load(); err != nil {
		log.Print("error cargando el archivo: ", err)
	}

	loglevel := os.Getenv("LOG")
	if loglevel == "" {
		config.LogLevel = "INFO"
	} else {
		config.LogLevel = loglevel
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
		return config, errors.New("missing DB url")
	}

	config.DbURL = dbUrl

	treatmentUrl := os.Getenv("TREATMENT_URL")
	if treatmentUrl == "" {
		return config, errors.New("missing treatment service url")

	}
	config.TreatmentURL = treatmentUrl
	return config, nil
}
