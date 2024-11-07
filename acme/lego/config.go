// Code from https://github.com/go-acme/lego/blob/v4.19.2/internal/dnsdocs/generator.go
// License: MIT

package lego

import (
	"acme-manager/logger"
	"acme-manager/secret"
	"acme-manager/util"
	"github.com/cockroachdb/errors"
)

type DnsProviderConfig struct {
	Credentials map[string]string `json:"credentials"`
	Additional  map[string]string `json:"additional,omitempty"`
}

func (config *DnsProviderConfig) Encrypt() error {
	for k, v := range config.Credentials {
		encryptV, err := secret.Get().Encrypt(v)
		if err != nil {
			return errors.Wrapf(err, "encrypt error for field: %s", k)
		}
		config.Credentials[k] = encryptV
	}
	return nil
}

func (config *DnsProviderConfig) Sensitive() {
	for k, v := range config.Credentials {
		if v == "encrypted" {
			continue
		}
		config.Credentials[k] = util.MakeSensitive(v)
	}
}

func (config *DnsProviderConfig) Decrypt() error {
	var decryptFailedFields []string

	for k, v := range config.Credentials {
		decryptV, err := secret.Get().Decrypt(v)
		if err != nil {
			logger.Warnf("decrypt error: %v", err)
			config.Credentials[k] = "encrypted"
			decryptFailedFields = append(decryptFailedFields, k)
		} else {
			config.Credentials[k] = decryptV
		}
	}
	if len(decryptFailedFields) > 0 {
		return errors.Errorf("decrypt failed for fields: %v", decryptFailedFields)
	}
	return nil
}
