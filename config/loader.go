package config

import "github.com/spf13/viper"

var AppConfig *Config

type Config struct {
	Port        string
	Environment string
	Content     string
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}
