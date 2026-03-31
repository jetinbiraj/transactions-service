package config

import (
	"log"
	"strings"
	"transactions-service/db"
	"transactions-service/server"

	"github.com/spf13/viper"
)

// registry is for the configuration values.
var registry *viper.Viper

// Set the configs
func Set() error {

	registry = viper.New()

	registry.SetConfigName("config")
	registry.SetConfigType("yaml")

	registry.AddConfigPath("./config")
	registry.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	registry.AutomaticEnv()

	if err := registry.ReadInConfig(); err != nil {
		log.Printf("Error reading settings file: %v", err)
		return err
	}
	return nil
}

// ServerConfig populates and returns all server configs
func ServerConfig() server.Config {
	return server.Config{
		Port: registry.GetString("SERVER_PORT"),
	}
}

// IsLogEnabled checks if logs are enabled by configuration
func IsLogEnabled() bool {
	return registry.GetBool("LOG_ENABLED")
}

// GetDBName gets the name of the db from config
func GetDBName() string {
	return registry.GetString("DB_NAME")
}

func GetPostgresConfig() db.PostgresConfig {
	return db.PostgresConfig{
		Host:            registry.GetString("POSTGRES_HOST"),
		Port:            registry.GetInt("POSTGRES_PORT"),
		User:            registry.GetString("POSTGRES_USER"),
		Password:        registry.GetString("POSTGRES_PASSWORD"),
		DBName:          registry.GetString("POSTGRES_DB_NAME"),
		SSLMode:         registry.GetString("POSTGRES_SSL_MODE"),
		MaxOpenConns:    registry.GetInt("POSTGRES_MAX_OPEN_CONNS"),
		MaxIdleConns:    registry.GetInt("POSTGRES_MAX_IDLE_CONNS"),
		ConnMaxLifetime: registry.GetDuration("POSTGRES_CONN_MAX_LIFETIME"),
	}
}
