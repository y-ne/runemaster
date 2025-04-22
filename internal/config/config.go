package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

func Load() (*Config, error) {
	// Set defaults
	viper.SetDefault("SERVER_PORT", "3000")
	viper.SetDefault("ENVIRONMENT", "development")
	viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/kurumi")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading .env file: %w", err)
		}
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &config, nil
}
