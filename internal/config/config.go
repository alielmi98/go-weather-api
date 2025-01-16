package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string
	WeatherAPIKey string
	RedisAddress  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
