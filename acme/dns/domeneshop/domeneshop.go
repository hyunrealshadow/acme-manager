// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/domeneshop
// License: MIT

package domeneshop

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"net/http"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/providers/dns/domeneshop"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvAPISecret          = "apiSecret"
	EnvAPIToken           = "apiToken"
	EnvHTTPTimeout        = "httpTimeout"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Domeneshop Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"apiToken\":{\"type\":\"string\",\"description\":\"API token\"},\"apiSecret\":{\"type\":\"string\",\"description\":\"API secret\"}},\"additional\":{\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("apiToken", "apiSecret")
var additionalFields = mapset.NewSet("pollingInterval", "propagationTimeout", "httpTimeout")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal domeneshop JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add domeneshop schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile domeneshop schema: %v", err)
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
		return nil, fmt.Errorf("domeneshop: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *domeneshop.Config {
	return &domeneshop.Config{
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, 5*time.Minute),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, 20*time.Second),
		HTTPClient: &http.Client{
			Timeout: env.GetOrDefaultSecond(EnvHTTPTimeout, 30*time.Second),
		},
	}
}

func NewDNSProvider(env lego.Env) (*domeneshop.DNSProvider, error) {
	values, err := env.Get(EnvAPIToken, EnvAPISecret)
	if err != nil {
		return nil, fmt.Errorf("domeneshop: %w", err)
	}

	config := NewConfig(env)
	config.APIToken = values[EnvAPIToken]
	config.APISecret = values[EnvAPISecret]

	return domeneshop.NewDNSProviderConfig(config)
}