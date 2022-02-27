package app

import (
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	Host           string
	Port           string
	Database       string
	User           string
	Password       string
	TimeoutSeconds int
}

type ApplicationConfig struct {
	Addr                      string
	DbConnectionsSettingsPath string
	DBConfig                  DBConfig
}

func parseConfig() *ApplicationConfig {
	// Server HTTP network address
	addr := getEnv("ADDR", ":8080")
	dbConnectionsSettingsPath := getEnv("DB_CONNECTIONS_SETTINGS_PATH", "./config/db_connections.json")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "pikachu")
	dbUser := getEnv("DB_USER", "pika")
	dbPassword := getEnv("DB_PASSWORD", "chu")
	dbTimeoutSeconds := getEnvInt("DB_TIMEOUT_SECONDS", 0)

	return &ApplicationConfig{
		Addr:                      addr,
		DbConnectionsSettingsPath: dbConnectionsSettingsPath,
		DBConfig: DBConfig{
			Host:           dbHost,
			Port:           dbPort,
			Database:       dbName,
			User:           dbUser,
			Password:       dbPassword,
			TimeoutSeconds: dbTimeoutSeconds,
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if valueString, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(valueString)
		if err == nil {
			return value
		}
		log.Printf("Failed to convert %s env value %s to int. Falling back to default %d", key, valueString, fallback)
	}

	return fallback
}
