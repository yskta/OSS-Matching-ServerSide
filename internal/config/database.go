package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDBConfig() (*DBConfig, error) {
	requiredEnvs := map[string]string{
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"DB_SSLMODE":  os.Getenv("DB_SSLMODE"),
	}

	for envName, envValue := range requiredEnvs {
		if envValue == "" {
			return nil, fmt.Errorf("required environment variable %s is not set", envName)
		}
	}

	return &DBConfig{
		Host:     requiredEnvs["DB_HOST"],
		Port:     requiredEnvs["DB_PORT"],
		User:     requiredEnvs["DB_USER"],
		Password: requiredEnvs["DB_PASSWORD"],
		DBName:   requiredEnvs["DB_NAME"],
		SSLMode:  requiredEnvs["DB_SSLMODE"],
	}, nil
}
