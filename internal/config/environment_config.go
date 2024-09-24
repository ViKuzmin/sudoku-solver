package config

import (
	"github.com/spf13/viper"
	"log"
)

type EnvironmentConfig struct {
	ServerConfig ServerConfig `yaml:"serverConfig"`
}

type ServerConfig struct {
	Env  string `yaml:"env" env-default:"local"`
	Port string `yaml:"port"`
}

func MustLoadEnvironmentConfig() *EnvironmentConfig {
	viper.SetConfigName("environment")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var cfg EnvironmentConfig

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
