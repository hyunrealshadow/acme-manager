// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/inwx
// License: MIT

package inwx

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/inwx"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvPassword           = "password"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
	EnvSandbox            = "sandbox"
	EnvSharedSecret       = "sharedSecret"
	EnvTTL                = "ttl"
	EnvUsername           = "username"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"INWX Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"username\":{\"type\":\"string\",\"description\":\"Username\"},\"password\":{\"type\":\"string\",\"description\":\"Password\"}},\"additional\":{\"sandbox\":{\"type\":\"string\",\"description\":\"Activate the sandbox (boolean)\"},\"sharedSecret\":{\"type\":\"string\",\"description\":\"shared secret related to 2FA\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation (default 360s)\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("username", "password")
var additionalFields = mapset.NewSet("sandbox", "sharedSecret", "pollingInterval", "propagationTimeout", "ttl")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal inwx JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add inwx schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile inwx schema: %v", err)
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
		return nil, fmt.Errorf("inwx: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *inwx.Config {
	return &inwx.Config{
		TTL: env.GetOrDefaultInt(EnvTTL, 300),
		// INWX has rather unstable propagation delays, thus using a larger default value
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, 360*time.Second),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, dns01.DefaultPollingInterval),
		Sandbox:            env.GetOrDefaultBool(EnvSandbox, false),
	}
}

func NewDNSProvider(env lego.Env) (*inwx.DNSProvider, error) {
	values, err := env.Get(EnvUsername, EnvPassword)
	if err != nil {
		return nil, fmt.Errorf("inwx: %w", err)
	}

	config := NewConfig(env)
	config.Username = values[EnvUsername]
	config.Password = values[EnvPassword]
	config.SharedSecret = env.GetOrFile(EnvSharedSecret)

	return inwx.NewDNSProviderConfig(config)
}