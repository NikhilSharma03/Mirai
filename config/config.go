package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	PORT string `mapstructure:"PORT"`
}

func NewConfig() (*Config, error) {
	environment := os.Getenv("GO_ENV")
	if environment == "production" {
		return &Config{
			PORT: os.Getenv("PORT"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
