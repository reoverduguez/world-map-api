package configs

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Database   Database
	ServerPort int `envconfig:"SERVER_PORT" default:"80"`
}

type Database struct {
	USERNAME string `envconfig:"DB_USERNAME" required:"true"`
	PASSWORD string `envconfig:"DB_PASSWORD" required:"true"`
}

func ParsedConfig() (Config, error) {
	database := Database{
		USERNAME: os.Getenv("DB_USERNAME"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
	}
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))

	if err != nil {
		err = fmt.Errorf("error converting port number to int %s", err)
		return Config{}, err
	}

	config := Config{
		Database:   database,
		ServerPort: serverPort,
	}

	return config, err
}
