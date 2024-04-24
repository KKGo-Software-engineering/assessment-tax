package config

import "os"

type Config struct {
	Server     *ServerConfig
	Auth       *AuthConfig
	PostgreSQL *PostgresConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type AuthConfig struct {
	Username string
	Password string
}

type PostgresConfig struct {
	URL string
}

func GetConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			Host: getEnv("HOST", "0.0.0.0"),
			Port: getEnv("PORT", "8080"),
		},
		Auth: &AuthConfig{
			Username: getEnv("ADMIN_USERNAME", ""),
			Password: getEnv("ADMIN_PASSWORD", ""),
		},
		PostgreSQL: &PostgresConfig{
			URL: getEnv("DATABASE_URL", ""),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
