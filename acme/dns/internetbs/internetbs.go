// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/internetbs
// License: MIT

package internetbs

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"net/http"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/internetbs"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvAPIKey             = "apiKey"
	EnvHTTPTimeout        = "httpTimeout"
	EnvPassword           = "password"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
	EnvTTL                = "ttl"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Internet.bs Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"apiKey\":{\"type\":\"string\",\"description\":\"API key\"},\"password\":{\"type\":\"string\",\"description\":\"API password\"}},\"additional\":{\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("apiKey", "password")
var additionalFields = mapset.NewSet("propagationTimeout", "ttl", "httpTimeout", "pollingInterval")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal internetbs JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add internetbs schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile internetbs schema: %v", err)
	}
}

func ParseConfig(conf *lego.DnsProviderConfig) (*lego.Env, error) {
	data := make(map[string]string)
	if conf.Credentials != nil {
		for key, value := range conf.Credentials {
			if credentialsFields.Contains(key) {
				data[key] = value
			} else {
				delete(conf.Credentials, key)
			}
		}
	}
	if conf.Additional != nil {
		for key, value := range conf.Additional {
			if additionalFields.Contains(key) {
				data[key] = value
			} else {
				delete(conf.Additional, key)
			}
		}
	}

	anyData := make(map[string]any)
	for key, value := range data {
		anyData[key] = value
	}
	if err := compiledJsonSchema.Validate(anyData); err != nil {
		return nil, fmt.Errorf("internetbs: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *internetbs.Config {
	return &internetbs.Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, 3600),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, dns01.DefaultPropagationTimeout),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, dns01.DefaultPollingInterval),
		HTTPClient: &http.Client{
			Timeout: env.GetOrDefaultSecond(EnvHTTPTimeout, 30*time.Second),
		},
	}
}

func NewDNSProvider(env lego.Env) (*internetbs.DNSProvider, error) {
	values, err := env.Get(EnvAPIKey, EnvPassword)
	if err != nil {
		return nil, fmt.Errorf("internetbs: %w", err)
	}

	config := NewConfig(env)
	config.APIKey = values[EnvAPIKey]
	config.Password = values[EnvPassword]

	return internetbs.NewDNSProviderConfig(config)
}
