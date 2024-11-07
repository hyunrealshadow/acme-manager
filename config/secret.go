package config

import (
	"github.com/spf13/viper"
)

var (
	secretProvider = "secret.provider"
	xchacha20Key   = "secret.xchacha20.key"
	vaultUrl       = "secret.vault.url"
	vaultToken     = "secret.vault.token"
)

const (
	XChaCha20 = "xchacha20"
	Vault     = "vault"
)

type XChaCha20Config struct {
	Key string
}

type VaultConfig struct {
	URL   string
	Token string
}

type SecretConfig struct {
	Provider  string
	XChaCha20 *XChaCha20Config
	Vault     *VaultConfig
}

func BindSecretEnv() {
	err := viper.BindEnv(secretProvider, "SECRET_PROVIDER")
	if err != nil {
		warnEnv(secretProvider, err)
	}
	err = viper.BindEnv(xchacha20Key, "CHACHA20_KEY")
	if err != nil {
		warnEnv(xchacha20Key, err)
	}
	err = viper.BindEnv(vaultUrl, "VAULT_URL")
	if err != nil {
		warnEnv(vaultUrl, err)
	}
	err = viper.BindEnv(vaultToken, "VAULT_TOKEN")
	if err != nil {
		warnEnv(vaultToken, err)
	}
}

func LoadSecretConfig() SecretConfig {
	if !viper.IsSet(secretProvider) {
		fatalKey(secretProvider)
	}
	var config = SecretConfig{
		Provider: viper.GetString(secretProvider),
	}
	if viper.IsSet(xchacha20Key) {
		config.XChaCha20 = &XChaCha20Config{
			Key: viper.GetString(xchacha20Key),
		}
	} else if viper.IsSet(vaultUrl) && viper.IsSet(vaultToken) {
		config.Vault = &VaultConfig{
			URL:   viper.GetString(vaultUrl),
			Token: viper.GetString(vaultToken),
		}
	} else {
		fatalKey(xchacha20Key + " or " + vaultUrl + " and " + vaultToken)
	}
	return config
}
