// Code from https://github.com/go-acme/lego/tree/v4.19.4/providers/dns/ovh
// License: MIT

package ovh

import (
	"acme-manager/acme/lego"
	"fmt"
	"github.com/go-acme/lego/v4/providers/dns/ovh"
)

func createConfigFromEnvVars(env lego.Env) (*ovh.Config, error) {
	firstAppKeyEnvVar := findFirstValuedEnvVar(env, EnvApplicationKey, EnvApplicationSecret, EnvConsumerKey)
	firstOAuth2EnvVar := findFirstValuedEnvVar(env, EnvClientID, EnvClientSecret)

	if firstAppKeyEnvVar != "" && firstOAuth2EnvVar != "" {
		return nil, fmt.Errorf("can't use both %s and %s at the same time", firstAppKeyEnvVar, firstOAuth2EnvVar)
	}

	config := NewConfig(env)

	if firstOAuth2EnvVar != "" {
		values, err := env.Get(EnvEndpoint, EnvClientID, EnvClientSecret)
		if err != nil {
			return nil, err
		}

		config.APIEndpoint = values[EnvEndpoint]
		config.OAuth2Config = &ovh.OAuth2Config{
			ClientID:     values[EnvClientID],
			ClientSecret: values[EnvClientSecret],
		}

		return config, nil
	}

	values, err := env.Get(EnvEndpoint, EnvApplicationKey, EnvApplicationSecret, EnvConsumerKey)
	if err != nil {
		return nil, err
	}

	config.APIEndpoint = values[EnvEndpoint]

	config.ApplicationKey = values[EnvApplicationKey]
	config.ApplicationSecret = values[EnvApplicationSecret]
	config.ConsumerKey = values[EnvConsumerKey]

	return config, nil
}

func findFirstValuedEnvVar(env lego.Env, envVars ...string) string {
	for _, envVar := range envVars {
		if env.GetOrFile(envVar) != "" {
			return envVar
		}
	}

	return ""
}
