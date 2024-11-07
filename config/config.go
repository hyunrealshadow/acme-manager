package config

import (
	"acme-manager/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Secret   SecretConfig
}

var config Config

func init() {
	viper.AddConfigPath("/etc/acme-manager")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()
	BindDatabaseEnv()
	BindSecretEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatalf("Failed read config: %v", err)
	}

	config = Config{
		Database: LoadDatabaseConfig(),
		Server:   LoadServerConfig(),
		Secret:   LoadSecretConfig(),
	}
}

func Get() Config {
	return config
}
