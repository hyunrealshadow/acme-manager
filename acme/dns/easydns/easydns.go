// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/easydns
// License: MIT

package easydns

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/easydns"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvEndpoint           = "endpoint"
	EnvHTTPTimeout        = "httpTimeout"
	EnvKey                = "key"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
	EnvSequenceInterval   = "sequenceInterval"
	EnvTTL                = "ttl"
	EnvToken              = "token"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"EasyDNS Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"token\":{\"type\":\"string\",\"description\":\"API Token\"},\"key\":{\"type\":\"string\",\"description\":\"API Key\"}},\"additional\":{\"sequenceInterval\":{\"type\":\"string\",\"description\":\"Time between sequential requests\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"},\"endpoint\":{\"type\":\"string\",\"description\":\"The endpoint URL of the API Server\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("token", "key")
var additionalFields = mapset.NewSet("sequenceInterval", "ttl", "httpTimeout", "endpoint", "pollingInterval", "propagationTimeout")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal easydns JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add easydns schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile easydns schema: %v", err)
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
		return nil, fmt.Errorf("easydns: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *easydns.Config {
	return &easydns.Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, dns01.DefaultTTL),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, dns01.DefaultPropagationTimeout),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, dns01.DefaultPollingInterval),
		SequenceInterval:   env.GetOrDefaultSecond(EnvSequenceInterval, dns01.DefaultPropagationTimeout),
		HTTPClient: &http.Client{
			Timeout: env.GetOrDefaultSecond(EnvHTTPTimeout, 30*time.Second),
		},
	}
}

func NewDNSProvider(env lego.Env) (*easydns.DNSProvider, error) {
	config := NewConfig(env)

	endpoint, err := url.Parse(env.GetOrDefaultString(EnvEndpoint, DefaultBaseURL))
	if err != nil {
		return nil, fmt.Errorf("easydns: %w", err)
	}
	config.Endpoint = endpoint

	values, err := env.Get(EnvToken, EnvKey)
	if err != nil {
		return nil, fmt.Errorf("easydns: %w", err)
	}

	config.Token = values[EnvToken]
	config.Key = values[EnvKey]

	return easydns.NewDNSProviderConfig(config)
}