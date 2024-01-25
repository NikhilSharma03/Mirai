package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	PORT              string `mapstructure:"PORT"`
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT     string `mapstructure:"POSTGRES_PORT"`
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`
}

func NewConfig() (*Config, error) {
	environment := os.Getenv("GO_ENV")
	if environment == "production" {
		return &Config{
			PORT:              os.Getenv("PORT"),
			POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
			POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
			POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
			POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
			POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
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
