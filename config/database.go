package config

import (
	"github.com/spf13/viper"
)

var databaseDSN = "database.dsn"

type DatabaseConfig struct {
	DSN string
}

func BindDatabaseEnv() {
	err := viper.BindEnv(databaseDSN, "DSN")
	if err != nil {
		warnEnv(databaseDSN, err)
	}
}

func LoadDatabaseConfig() DatabaseConfig {
	if !viper.IsSet(databaseDSN) {
		fatalKey(databaseDSN)
	}
	return DatabaseConfig{
		DSN: viper.GetString(databaseDSN),
	}
}
