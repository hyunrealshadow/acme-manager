// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/alidns
// License: MIT

package alidns

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvAccessKey          = "accessKey"
	EnvHTTPTimeout        = "httpTimeout"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
	EnvRAMRole            = "ramRole"
	EnvRegionID           = "regionId"
	EnvSecretKey          = "secretKey"
	EnvSecurityToken      = "securityToken"
	EnvTTL                = "ttl"
	defaultRegionID       = "cn-hangzhou"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Alibaba Cloud DNS Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"accessKey\":{\"type\":\"string\",\"description\":\"Access key ID\"},\"secretKey\":{\"type\":\"string\",\"description\":\"Access Key secret\"},\"securityToken\":{\"type\":\"string\",\"description\":\"STS Security Token (optional)\"},\"ramRole\":{\"type\":\"string\",\"description\":\"Your instance RAM role (https://www.alibabacloud.com/help/doc-detail/54579.htm)\"}},\"additional\":{\"regionId\":{\"type\":\"string\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("accessKey", "secretKey", "securityToken", "ramRole")
var additionalFields = mapset.NewSet("regionId", "pollingInterval", "propagationTimeout", "ttl", "httpTimeout")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal alidns JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add alidns schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile alidns schema: %v", err)
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
		return nil, fmt.Errorf("alidns: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *alidns.Config {
	return &alidns.Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, 600),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, dns01.DefaultPropagationTimeout),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, dns01.DefaultPollingInterval),
		HTTPTimeout:        env.GetOrDefaultSecond(EnvHTTPTimeout, 10*time.Second),
	}
}

func NewDNSProvider(env lego.Env) (*alidns.DNSProvider, error) {
	config := NewConfig(env)
	config.RegionID = env.GetOrFile(EnvRegionID)

	values, err := env.Get(EnvRAMRole)
	if err == nil {
		config.RAMRole = values[EnvRAMRole]
		return alidns.NewDNSProviderConfig(config)
	}

	values, err = env.Get(EnvAccessKey, EnvSecretKey)
	if err != nil {
		return nil, fmt.Errorf("alicloud: %w", err)
	}

	config.APIKey = values[EnvAccessKey]
	config.SecretKey = values[EnvSecretKey]
	config.SecurityToken = env.GetOrFile(EnvSecurityToken)

	return alidns.NewDNSProviderConfig(config)
}
