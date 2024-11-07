package config

import "github.com/spf13/viper"

var serverPort = "server.port"
var serverPlayground = "server.playground"

type ServerConfig struct {
	Port       uint16
	Playground bool
}

func BindServerEnv() {
	err := viper.BindEnv(serverPort, "SERVER_PORT")
	if err != nil {
		warnEnv(serverPort, err)
	}
	err = viper.BindEnv(serverPlayground, "SERVER_PLAYGROUND")
	if err != nil {
		warnEnv(serverPlayground, err)
	}
}

func LoadServerConfig() ServerConfig {
	if !viper.IsSet(serverPort) {
		fatalKey(serverPort)
	}
	return ServerConfig{
		Port:       viper.GetUint16(serverPort),
		Playground: viper.GetBool(serverPlayground),
	}
}
