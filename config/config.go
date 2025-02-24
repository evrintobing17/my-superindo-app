package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB     DatabaseConfig
	Server ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshalling config: %v", err)
		return nil, err
	}

	return &config, nil
}
