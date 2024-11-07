package config

import (
	"acme-manager/logger"
)

func warnEnv(key string, err error) {
	logger.Warnf("Error binding env var %s : %s \n", key, err)
}

func fatalKey(key string) {
	logger.Fatalf("Config %s is required", key)
}
