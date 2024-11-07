package secret

import (
	"acme-manager/config"
	"acme-manager/logger"
)

type Provider interface {
	Encrypt(data string) (string, error)
	Decrypt(data string) (string, error)
}

var provider Provider

func init() {
	var cfg = config.Get()
	switch cfg.Secret.Provider {
	case config.XChaCha20:
		provider = NewXChaCha2Provider(cfg.Secret.XChaCha20.Key)
	default:
		logger.Fatalf("Unsupported secret provider: %s", cfg.Secret.Provider)
	}
}

func Get() Provider {
	return provider
}
