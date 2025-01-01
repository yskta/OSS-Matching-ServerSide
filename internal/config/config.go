package config

import "fmt"

type Config struct {
	DB *DBConfig
}

func NewConfig() (*Config, error) {
	dbConfig, err := NewDBConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %w", err)
	}
	
	return &Config{
		DB: dbConfig,
	}, nil
}
