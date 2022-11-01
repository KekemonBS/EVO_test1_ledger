package env

import (
	"fmt"
	"os"
)

type Config struct {
	PostgresURI string
}

func NewConfig() (*Config, error) {
	// postgresql://[userspec@][hostspec][/dbname][?paramspec]
	postgresURI, ok := os.LookupEnv("POSTGRESURI")
	if !ok {
		return nil, fmt.Errorf("no POSTGRESURI env variable")
	}
	return &Config{
		PostgresURI: postgresURI,
	}, nil
}
