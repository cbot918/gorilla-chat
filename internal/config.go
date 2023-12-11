package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBType   string
	Username string
	Password string
	Hostname string
	Database string
}

func NewConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DBType:   os.Getenv("DBType"),
		Username: os.Getenv("Username"),
		Password: os.Getenv("Password"),
		Hostname: os.Getenv("Hostname"),
		Database: os.Getenv("Database"),
	}, nil
}
