// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/webnames
// License: MIT

package webnames

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"net/http"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/webnames"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvAPIKey             = "apiKey"
	EnvHTTPTimeout        = "httpTimeout"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Webnames Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"apiKey\":{\"type\":\"string\",\"description\":\"Domain API key\"}},\"additional\":{\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("apiKey")
var additionalFields = mapset.NewSet("ttl", "httpTimeout", "pollingInterval", "propagationTimeout")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal webnames JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add webnames schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile webnames schema: %v", err)
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
		return nil, fmt.Errorf("webnames: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *webnames.Config {
	return &webnames.Config{
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, dns01.DefaultPropagationTimeout),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, dns01.DefaultPollingInterval),
		HTTPClient: &http.Client{
			Timeout: env.GetOrDefaultSecond(EnvHTTPTimeout, 30*time.Second),
		},
	}
}

func NewDNSProvider(env lego.Env) (*webnames.DNSProvider, error) {
	values, err := env.Get(EnvAPIKey)
	if err != nil {
		return nil, fmt.Errorf("webnames: %w", err)
	}

	config := NewConfig(env)
	config.APIKey = values[EnvAPIKey]

	return webnames.NewDNSProviderConfig(config)
}
