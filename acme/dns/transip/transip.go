// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/transip
// License: MIT

package transip

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/providers/dns/transip"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvAccountName        = "accountName"
	EnvPollingInterval    = "pollingInterval"
	EnvPrivateKeyPath     = "privateKeyPath"
	EnvPropagationTimeout = "propagationTimeout"
	EnvTTL                = "ttl"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"TransIP Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"accountName\":{\"type\":\"string\",\"description\":\"Account name\"},\"privateKeyPath\":{\"type\":\"string\",\"description\":\"Private key path\"}},\"additional\":{\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("accountName", "privateKeyPath")
var additionalFields = mapset.NewSet("ttl", "pollingInterval", "propagationTimeout")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal transip JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add transip schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile transip schema: %v", err)
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
		return nil, fmt.Errorf("transip: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *transip.Config {
	return &transip.Config{
		TTL:                int64(env.GetOrDefaultInt(EnvTTL, 10)),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, 10*time.Minute),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, 10*time.Second),
	}
}

func NewDNSProvider(env lego.Env) (*transip.DNSProvider, error) {
	values, err := env.Get(EnvAccountName, EnvPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("transip: %w", err)
	}

	config := NewConfig(env)
	config.AccountName = values[EnvAccountName]
	config.PrivateKeyPath = values[EnvPrivateKeyPath]

	return transip.NewDNSProviderConfig(config)
}
