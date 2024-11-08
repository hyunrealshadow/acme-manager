// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/scaleway
// License: MIT

package scaleway

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/providers/dns/scaleway"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvAPIToken               = "apiToken"
	EnvAccessKey              = "accessKey"
	EnvPollingInterval        = "pollingInterval"
	EnvProjectID              = "projectId"
	EnvPropagationTimeout     = "propagationTimeout"
	EnvSecretKey              = "secretKey"
	EnvTTL                    = "ttl"
	defaultPollingInterval    = 10000000000
	defaultPropagationTimeout = 120000000000
	dumpAccessKey             = "SCWXXXXXXXXXXXXXXXXX"
	minTTL                    = 60
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Scaleway Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"secretKey\":{\"type\":\"string\",\"description\":\"Secret key\"},\"projectId\":{\"type\":\"string\",\"description\":\"Project to use (optional)\"}},\"additional\":{\"apiToken\":{\"type\":\"string\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"accessKey\":{\"type\":\"string\",\"description\":\"Access key\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("secretKey", "projectId")
var additionalFields = mapset.NewSet("apiToken", "propagationTimeout", "ttl", "accessKey", "pollingInterval")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal scaleway JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add scaleway schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile scaleway schema: %v", err)
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
		return nil, fmt.Errorf("scaleway: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(e lego.Env) *scaleway.Config {
	return &scaleway.Config{
		AccessKey:          dumpAccessKey,
		TTL:                lego.GetOneWithFallback(e, EnvTTL, minTTL, strconv.Atoi, altEnvName(EnvTTL)),
		PropagationTimeout: lego.GetOneWithFallback(e, EnvPropagationTimeout, defaultPropagationTimeout, lego.ParseSecond, altEnvName(EnvPropagationTimeout)),
		PollingInterval:    lego.GetOneWithFallback(e, EnvPollingInterval, defaultPollingInterval, lego.ParseSecond, altEnvName(EnvPollingInterval)),
	}
}

func NewDNSProvider(env lego.Env) (*scaleway.DNSProvider, error) {
	values, err := env.GetWithFallback([]string{EnvSecretKey, EnvAPIToken})
	if err != nil {
		return nil, fmt.Errorf("scaleway: %w", err)
	}

	config := NewConfig(env)
	config.Token = values[EnvSecretKey]
	config.AccessKey = env.GetOrDefaultString(EnvAccessKey, dumpAccessKey)
	config.ProjectID = env.GetOrFile(EnvProjectID)

	return scaleway.NewDNSProviderConfig(config)
}
