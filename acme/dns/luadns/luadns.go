// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/luadns
// License: MIT

package luadns

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"net/http"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/providers/dns/luadns"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvAPIToken           = "apiToken"
	EnvAPIUsername        = "apiUsername"
	EnvHTTPTimeout        = "httpTimeout"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
	EnvTTL                = "ttl"
	minTTL                = 300
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"LuaDNS Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"apiUsername\":{\"type\":\"string\",\"description\":\"Username (your email)\"},\"apiToken\":{\"type\":\"string\",\"description\":\"API token\"}},\"additional\":{\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("apiUsername", "apiToken")
var additionalFields = mapset.NewSet("pollingInterval", "propagationTimeout", "ttl", "httpTimeout")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal luadns JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add luadns schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile luadns schema: %v", err)
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
		return nil, fmt.Errorf("luadns: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *luadns.Config {
	return &luadns.Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, minTTL),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, 120*time.Second),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, 2*time.Second),
		HTTPClient: &http.Client{
			Timeout: env.GetOrDefaultSecond(EnvHTTPTimeout, 30*time.Second),
		},
	}
}

func NewDNSProvider(env lego.Env) (*luadns.DNSProvider, error) {
	values, err := env.Get(EnvAPIUsername, EnvAPIToken)
	if err != nil {
		return nil, fmt.Errorf("luadns: %w", err)
	}

	config := NewConfig(env)
	config.APIUsername = values[EnvAPIUsername]
	config.APIToken = values[EnvAPIToken]

	return luadns.NewDNSProviderConfig(config)
}