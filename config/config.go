package config

import (
	"os"
	"strconv"
)

// Estructura de SQL
type SQLConfig struct {
	Database string
	Username string
	Password string
	Host     string
	Port     int
}

// Estructura principal
type Config struct {
	SQL SQLConfig
}

// Retorna configuracion
func New() *Config {
	return &Config{
		SQL: SQLConfig{
			Database: getEnv("DB_NAME", ""),
			Username: getEnv("DB_USER", ""),
			Password: getEnv("DB_PASS", ""),
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnvAsInt("DB_PORT", 0),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
