// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/cloudns
// License: MIT

package cloudns

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"net/http"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/providers/dns/cloudns"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvAuthID             = "authId"
	EnvAuthPassword       = "authPassword"
	EnvHTTPTimeout        = "httpTimeout"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
	EnvSubAuthID          = "subAuthId"
	EnvTTL                = "ttl"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"ClouDNS Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"authId\":{\"type\":\"string\",\"description\":\"The API user ID\"},\"authPassword\":{\"type\":\"string\",\"description\":\"The password for API user ID\"}},\"additional\":{\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"},\"subAuthId\":{\"type\":\"string\",\"description\":\"The API sub user ID\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("authId", "authPassword")
var additionalFields = mapset.NewSet("httpTimeout", "subAuthId", "pollingInterval", "propagationTimeout", "ttl")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal cloudns JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add cloudns schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile cloudns schema: %v", err)
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
		return nil, fmt.Errorf("cloudns: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *cloudns.Config {
	return &cloudns.Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, 60),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, 180*time.Second),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, 10*time.Second),
		HTTPClient: &http.Client{
			Timeout: env.GetOrDefaultSecond(EnvHTTPTimeout, 30*time.Second),
		},
	}
}

func NewDNSProvider(env lego.Env) (*cloudns.DNSProvider, error) {
	var subAuthID string
	authID := env.GetOrFile(EnvAuthID)
	if authID == "" {
		subAuthID = env.GetOrFile(EnvSubAuthID)
	}

	if authID == "" && subAuthID == "" {
		return nil, fmt.Errorf("ClouDNS: some credentials information are missing: %s or %s", EnvAuthID, EnvSubAuthID)
	}

	values, err := env.Get(EnvAuthPassword)
	if err != nil {
		return nil, fmt.Errorf("ClouDNS: %w", err)
	}

	config := NewConfig(env)
	config.AuthID = authID
	config.SubAuthID = subAuthID
	config.AuthPassword = values[EnvAuthPassword]

	return cloudns.NewDNSProviderConfig(config)
}
