// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/ultradns
// License: MIT

package ultradns

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/providers/dns/ultradns"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvEndpoint           = "endpoint"
	EnvPassword           = "password"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
	EnvTTL                = "ttl"
	EnvUsername           = "username"
	defaultEndpoint       = "https://api.ultradns.com/"
	defaultUserAgent      = "go-acme/lego"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Ultradns Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"username\":{\"type\":\"string\",\"description\":\"API Username\"},\"password\":{\"type\":\"string\",\"description\":\"API Password\"}},\"additional\":{\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"endpoint\":{\"type\":\"string\",\"description\":\"API endpoint URL, defaults to https://api.ultradns.com/\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("username", "password")
var additionalFields = mapset.NewSet("propagationTimeout", "endpoint", "ttl", "pollingInterval")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal ultradns JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add ultradns schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile ultradns schema: %v", err)
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
		return nil, fmt.Errorf("ultradns: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *ultradns.Config {
	return &ultradns.Config{
		Endpoint:           env.GetOrDefaultString(EnvEndpoint, defaultEndpoint),
		TTL:                env.GetOrDefaultInt(EnvTTL, 120),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, 2*time.Minute),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, 4*time.Second),
	}
}

func NewDNSProvider(env lego.Env) (*ultradns.DNSProvider, error) {
	values, err := env.Get(EnvUsername, EnvPassword)
	if err != nil {
		return nil, fmt.Errorf("ultradns: %w", err)
	}

	config := NewConfig(env)
	config.Username = values[EnvUsername]
	config.Password = values[EnvPassword]

	return ultradns.NewDNSProviderConfig(config)
}