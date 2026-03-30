package config

import (
	"log"
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
