// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/bluecat
// License: MIT

package bluecat

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"net/http"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/bluecat"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvConfigName         = "configName"
	EnvDNSView            = "dnsView"
	EnvDebug              = "debug"
	EnvHTTPTimeout        = "httpTimeout"
	EnvPassword           = "password"
	EnvPollingInterval    = "pollingInterval"
	EnvPropagationTimeout = "propagationTimeout"
	EnvServerURL          = "serverUrl"
	EnvSkipDeploy         = "skipDeploy"
	EnvTTL                = "ttl"
	EnvUserName           = "userName"
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Bluecat Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"userName\":{\"type\":\"string\",\"description\":\"API username\"},\"password\":{\"type\":\"string\",\"description\":\"API password\"},\"configName\":{\"type\":\"string\",\"description\":\"Configuration name\"},\"dnsView\":{\"type\":\"string\",\"description\":\"External DNS View Name\"},\"serverUrl\":{\"type\":\"string\",\"description\":\"The server URL, should have scheme, hostname, and port (if required) of the authoritative Bluecat BAM serve\"}},\"additional\":{\"debug\":{\"type\":\"string\"},\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"},\"skipDeploy\":{\"type\":\"string\",\"description\":\"Skip deployements\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"},\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("userName", "password", "configName", "dnsView", "serverUrl")
var additionalFields = mapset.NewSet("debug", "httpTimeout", "skipDeploy", "pollingInterval", "propagationTimeout", "ttl")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal bluecat JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add bluecat schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile bluecat schema: %v", err)
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
		return nil, fmt.Errorf("bluecat: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *bluecat.Config {
	return &bluecat.Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, dns01.DefaultTTL),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, dns01.DefaultPropagationTimeout),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, dns01.DefaultPollingInterval),
		HTTPClient: &http.Client{
			Timeout: env.GetOrDefaultSecond(EnvHTTPTimeout, 30*time.Second),
		},
		Debug:      env.GetOrDefaultBool(EnvDebug, false),
		SkipDeploy: env.GetOrDefaultBool(EnvSkipDeploy, false),
	}
}

func NewDNSProvider(env lego.Env) (*bluecat.DNSProvider, error) {
	values, err := env.Get(EnvServerURL, EnvUserName, EnvPassword, EnvConfigName, EnvDNSView)
	if err != nil {
		return nil, fmt.Errorf("bluecat: %w", err)
	}

	config := NewConfig(env)
	config.BaseURL = values[EnvServerURL]
	config.UserName = values[EnvUserName]
	config.Password = values[EnvPassword]
	config.ConfigName = values[EnvConfigName]
	config.DNSView = values[EnvDNSView]

	return bluecat.NewDNSProviderConfig(config)
}